package appdeploy

import (
	"fmt"
	"maps"
	"path"
	"slices"
	"strings"

	"github.com/medvesek/infrastructure/lkw/assets"
	"github.com/medvesek/infrastructure/lkw/lib/cmd"
	"github.com/medvesek/infrastructure/lkw/lib/utils"
	"github.com/medvesek/infrastructure/lkw/src/deploy"
	"github.com/medvesek/infrastructure/lkw/src/dns"
	"github.com/medvesek/infrastructure/lkw/src/remote"
	"github.com/spf13/viper"
)

type AppDeploy struct{}

type Config struct {
	Source string
	Name   string
	Expose []struct {
		Service string
		Expose  []struct {
			Domain     string
			Port       int
			Entrypoint string
		}
	}
	PreDeploy []struct {
		Name    string
		Image   string
		Command string
		EnvFile string
	}
}

func New() *AppDeploy {
	return &AppDeploy{}
}

func (ad *AppDeploy) Run(config Config) error {
	name := config.Name
	serviceName := utils.DockerName(name)

	domains := getDomainsFromConfig(config)
	dns.SetupDomains(domains)

	tempDir, cleanup, err := deploy.PrepareSupportFiles(getTemplateItems(config), assets.NoFiles)
	defer cleanup()
	if err != nil {
		return err
	}

	source := path.Join(path.Dir(viper.GetString("config")), config.Source)

	cmd.Rsync(source+"/", tempDir+"/")

	remoteClient := remote.New()

	remoteDestination := remote.GetDestination(name)

	deployCommand := fmt.Sprintf(
		"cd %[1]s; dotenv docker stack deploy --compose-file %[1]s/docker-compose.yaml --compose-file %[1]s/override.yaml --detach=false --prune %[2]s",
		remoteDestination,
		serviceName,
	)

	if err := remoteClient.EnsureDir(remoteDestination); err != nil {
		return err
	}
	if err := remoteClient.RsyncD(tempDir+"/", remoteDestination); err != nil {
		return err
	}

	isFirstDeploy := false
	if err := remoteClient.Cmd("docker stack ps " + serviceName); err != nil {
		isFirstDeploy = true
	}

	if isFirstDeploy {
		if err := remoteClient.Cmd(deployCommand); err != nil {
			return err
		}
	}

	for _, preCommand := range createPreCommands(config, remoteDestination) {
		err := remoteClient.Cmd(preCommand)
		if err != nil {
			return err
		}
	}

	if !isFirstDeploy {
		if err := remoteClient.Cmd(deployCommand); err != nil {
			return err
		}
	}

	for _, domain := range domains {
		fmt.Println("https://" + domain)
	}

	return nil
}

func (ad *AppDeploy) Remove(config Config) {
	domains := getDomainsFromConfig(config)
	dns.RemoveDomains(domains)

	name := config.Name
	serviceName := utils.DockerName(name)

	remoteClient := remote.New()

	destination := remote.GetDestination(name)

	remoteClient.Cmd("docker stack rm  " + serviceName)

	remoteClient.Cmd("rm -rf " + destination)
}

func getTemplateItems(config Config) []deploy.TemplateItem {
	return []deploy.TemplateItem{
		{
			String: assets.AppDockerComposeOverrideTemplate,
			Data: assets.AppDockerComposeOverrideData{
				Name:   utils.DockerName(config.Name),
				Expose: config.Expose,
			},
			FileName: "override.yaml",
		},
	}
}

func getDomainsFromConfig(config Config) []string {
	domainsMap := map[string]bool{}
	for _, exposedService := range config.Expose {
		for _, exposedConfig := range exposedService.Expose {
			domainsMap[exposedConfig.Domain] = true
		}
	}
	return slices.Collect(maps.Keys(domainsMap))
}

func createPreCommands(config Config, remoteDestination string) []string {
	var preCommands []string
	for _, preCommandConfig := range config.PreDeploy {
		var preCommand strings.Builder
		preCommand.WriteString("docker run ")
		preCommand.WriteString("--rm ")
		preCommand.WriteString(fmt.Sprintf("--network %s_default ", utils.DockerName(config.Name)))
		preCommand.WriteString("--pull=always ")
		if preCommandConfig.EnvFile != "" {
			fmt.Fprintf(&preCommand, "--env-file %s/%s ", remoteDestination, preCommandConfig.EnvFile)
		}
		preCommand.WriteString(preCommandConfig.Image)
		preCommand.WriteString(" ")
		preCommand.WriteString(preCommandConfig.Command)
		preCommands = append(preCommands, preCommand.String())
	}
	return preCommands
}

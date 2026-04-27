package staticdeploy

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/assets"
	"github.com/medvesek/infrastructure/lkw/lib/cmd"
	"github.com/medvesek/infrastructure/lkw/lib/utils"
	"github.com/medvesek/infrastructure/lkw/src/deploy"
	"github.com/medvesek/infrastructure/lkw/src/dns"
	"github.com/medvesek/infrastructure/lkw/src/remote"
)

type StaticDeploy struct{}

type Config struct {
	Domain string
	Source string
	Spa    bool
}

func New() *StaticDeploy {
	return &StaticDeploy{}
}

func (sd *StaticDeploy) Run(config Config) error {
	name := utils.DockerName(config.Domain)
	err := dns.SetupDomain(config.Domain)

	if err != nil {
		return err
	}

	tempDir, cleanup, err := deploy.PrepareSupportFiles(getTemplateItems(config), assets.StaticFiles)
	defer cleanup()
	if err != nil {
		return err
	}

	cmd.Rsync(utils.EnsureTrailingSlash(config.Source), tempDir+"/public/")

	remoteClient := remote.New()

	destination := remote.GetDestination(config.Domain)

	if err := remoteClient.EnsureDir(destination); err != nil {
		return err
	}
	if err := remoteClient.RsyncD(tempDir+"/", destination); err != nil {
		return err
	}
	deployCommand := fmt.Sprintf("docker stack deploy --compose-file %s/docker-compose.yaml --detach=true %s", destination, name)
	if err := remoteClient.Cmd(deployCommand); err != nil {
		return err
	}

	fmt.Println("https://" + config.Domain)

	return nil
}

func getTemplateItems(config Config) []deploy.TemplateItem {
	name := utils.DockerName(config.Domain)
	return []deploy.TemplateItem{
		{
			String:   assets.StaticDockerComposeTemplate,
			Data:     assets.StaticDockerComposeData{Name: name, Domain: config.Domain},
			FileName: "docker-compose.yaml",
		},
		{
			String:   assets.StaticNginxDefaultConfTemplate,
			Data:     assets.StaticNginxDefaultConfData{Spa: config.Spa},
			FileName: "default.conf",
		},
	}
}

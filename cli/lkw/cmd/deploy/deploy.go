package deploy

import (
	"fmt"
	"slices"

	"github.com/medvesek/infrastructure/lkw/assets"
	"github.com/medvesek/infrastructure/lkw/constants"
	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/medvesek/infrastructure/lkw/lib/deploy"
	"github.com/medvesek/infrastructure/lkw/lib/remote"
	"github.com/medvesek/infrastructure/lkw/lib/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validTypes = []string{"static", "spa"}

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy to app server",
	Long:  "Deploy to app server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("domain") == "" {
			return fmt.Errorf("domain is required")
		}
		if viper.GetString("source") == "" {
			return fmt.Errorf("source is required")
		}
		if !slices.Contains(validTypes, viper.GetString("type")) {
			return fmt.Errorf("type must be one of: %s", validTypes)
		}
		return runDeploy()
	},
}

func init() {
	DeployCmd.Flags().String("type", "", "Type of the deploy [static|spa]")
	DeployCmd.Flags().String("domain", "", "The domain e.g. mysite.domain.com")
	DeployCmd.Flags().String("source", "", "Directory to deploy")

	viper.BindPFlag("type", DeployCmd.Flags().Lookup("type"))
	viper.BindPFlag("domain", DeployCmd.Flags().Lookup("domain"))
	viper.BindPFlag("source", DeployCmd.Flags().Lookup("source"))
}

func runDeploy() error {
	ip := viper.GetString("server_ip")
	user := viper.GetString("ssh_user")
	domain := viper.GetString("domain")
	source := viper.GetString("source")
	deployType := viper.GetString("type")

	name := utils.NameFromDomain(domain)

	err := setupDomain(domain, ip)

	if err != nil {
		return err
	}

	templates := []deploy.TemplateItem{
		{
			String:   assets.DockerComposeTemplate,
			Data:     assets.DockerComposeData{Name: name, Domain: domain},
			FileName: "docker-compose.yaml",
		},
		{
			String:   assets.NginxDefaultConf,
			Data:     assets.NginxDefaultConfData{Type: deployType},
			FileName: "default.conf",
		},
	}

	supportFilesDir, cleanup, err := deploy.PrepareSupportFiles(templates, assets.StaticFiles)
	defer cleanup()
	if err != nil {
		return err
	}

	remoteClient := remote.NewRemoteClient(user, ip)

	destination := fmt.Sprintf("~/%s/%s", constants.RemoteDir, domain)

	if err := remoteClient.Cmd(fmt.Sprintf("mkdir -p ~/%s", constants.RemoteDir)); err != nil {
		return err
	}
	if err := remoteClient.Rsync(supportFilesDir+"/", destination); err != nil {
		return err
	}
	if err := remoteClient.Rsync(source+"/", destination+"/public"); err != nil {
		return err
	}
	if err := remoteClient.Cmd(fmt.Sprintf("docker stack deploy --compose-file %s/docker-compose.yaml --detach=true %s", destination, name)); err != nil {
		return err
	}

	return nil
}

func setupDomain(domain string, ip string) error {
	cloudflareToken := viper.GetString("cloudflare_token")

	cloudflareClient := cloudflare.NewCloudflareClient(cloudflareToken)

	_, err := cloudflareClient.EnsureARecord(domain, ip)

	return err
}

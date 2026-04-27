package rm

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/medvesek/infrastructure/lkw/lib/utils"
	"github.com/medvesek/infrastructure/lkw/src/constants"
	"github.com/medvesek/infrastructure/lkw/src/remote"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var domain string
var name string

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove app",
	Long:  "Remove app",
	RunE: func(cmd *cobra.Command, args []string) error {
		return remove()
	},
}

func init() {
	RmCmd.Flags().StringVar(&domain, "domain", "", "The domain e.g. mysite.domain.com")
	RmCmd.Flags().StringVar(&name, "name", "", "The name of the app e.g. mysite.domain.com")

	RmCmd.MarkFlagRequired("domain")
}

func remove() error {
	cloudflareToken := viper.GetString("cloudflare_token")

	cloudflareClient := cloudflare.New(cloudflareToken)

	remoteClient := remote.New()

	if name == "" {
		name = domain
	}

	remoteClient.Cmd(fmt.Sprintf("docker stack rm %s", utils.DockerName(name)))
	remoteClient.Cmd(fmt.Sprintf("rm -rf ~/%s/%s", constants.RemoteDir, name))

	_, err := cloudflareClient.RemoveARecord(domain)

	if err != nil {
		return err
	}

	return nil
}

package remove

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/constants"
	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/medvesek/infrastructure/lkw/lib/remote"
	"github.com/medvesek/infrastructure/lkw/lib/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var domain string

var RemoveCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove app",
	Long:  "Remove app",
	RunE: func(cmd *cobra.Command, args []string) error {
		return remove()
	},
}

func init() {
	RemoveCmd.Flags().StringVar(&domain, "domain", "", "The domain e.g. mysite.domain.com")

	RemoveCmd.MarkFlagRequired("domain")
}

func remove() error {
	ip := viper.GetString("server_ip")
	user := viper.GetString("ssh_user")
	cloudflareToken := viper.GetString("cloudflare_token")

	cloudflareClient := cloudflare.NewCloudflareClient(cloudflareToken)

	remoteClient := remote.NewRemoteClient(user, ip)

	name := utils.NameFromDomain(domain)

	remoteClient.Cmd(fmt.Sprintf("docker stack rm %s", name))
	remoteClient.Cmd(fmt.Sprintf("rm -rf ~/%s/%s", constants.RemoteDir, domain))

	_, err := cloudflareClient.RemoveARecord(domain)

	if err != nil {
		return err
	}

	return nil
}

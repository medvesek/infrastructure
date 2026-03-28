package ensure

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var EnsureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "Ensure A record is present",
	Long:  "Ensure A record is present",
	Run: func(cmd *cobra.Command, args []string) {
		ensureARecord()
	},
}

var domain string
var ip string

func init() {
	EnsureCmd.Flags().StringVar(&domain, "domain", "", "The domain e.g. mysite.domain.com")
	EnsureCmd.Flags().StringVar(&ip, "ip", "", "The ip of the app server")
	EnsureCmd.MarkFlagRequired("domain")
	EnsureCmd.MarkFlagRequired("ip")
}

func ensureARecord() {
	cloudflareToken := viper.GetString("cloudflare_token")

	client := cloudflare.NewCloudflareClient(cloudflareToken)

	_, err := client.EnsureARecord(domain, ip)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Record present")
}

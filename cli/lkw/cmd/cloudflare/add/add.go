package add

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new A record",
	Long:  "Add a new A record",
	Run: func(cmd *cobra.Command, args []string) {
		addARecord()
	},
}

var domain string
var ip string

func init() {
	AddCmd.Flags().StringVar(&domain, "domain", "", "The domain e.g. mysite.domain.com")
	AddCmd.Flags().StringVar(&ip, "ip", "", "The ip of the app server")
	AddCmd.MarkFlagRequired("domain")
	AddCmd.MarkFlagRequired("ip")
}

func addARecord() {
	cloudflareToken := viper.GetString("cloudflare_token")

	client := cloudflare.NewCloudflareClient(cloudflareToken)

	_, err := client.CreateARecord(domain, ip)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Record added")
}

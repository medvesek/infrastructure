package remove

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an A record",
	Long:  "Remove an A record",
	Run: func(cmd *cobra.Command, args []string) {
		removeARecord()
	},
}

var domain string

func init() {
	RemoveCmd.Flags().StringVar(&domain, "domain", "", "The domain e.g. mysite.domain.com")
	RemoveCmd.MarkFlagRequired("domain")
}

func removeARecord() {
	cloudflareToken := viper.GetString("cloudflare_token")

	client := cloudflare.NewCloudflareClient(cloudflareToken)

	_, err := client.RemoveARecord(domain)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Record removed")
}

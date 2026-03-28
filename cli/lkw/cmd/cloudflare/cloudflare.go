package cloudflare

import (
	"github.com/medvesek/infrastructure/lkw/cmd/cloudflare/add"
	"github.com/medvesek/infrastructure/lkw/cmd/cloudflare/ensure"
	"github.com/medvesek/infrastructure/lkw/cmd/cloudflare/remove"
	"github.com/spf13/cobra"
)

var CloudflareCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "Cloudflare commands",
	Long:  "Commands for interacting with Cloudflare API",
}

func init() {
	CloudflareCmd.AddCommand(add.AddCmd)
	CloudflareCmd.AddCommand(remove.RemoveCmd)
	CloudflareCmd.AddCommand(ensure.EnsureCmd)
}

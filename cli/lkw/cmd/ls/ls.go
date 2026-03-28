package ls

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/constants"
	"github.com/medvesek/infrastructure/lkw/lib/remote"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List deployed apps",
	Long:  "List deployed apps",
	Run: func(cmd *cobra.Command, args []string) {
		ls()
	},
}

func ls() {
	ip := viper.GetString("server_ip")
	user := viper.GetString("ssh_user")
	dir := "~/" + constants.RemoteDir

	remoteClient := remote.NewRemoteClient(user, ip)

	remoteClient.Cmd(fmt.Sprintf("ls %s", dir))

}

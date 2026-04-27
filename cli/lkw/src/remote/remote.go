package remote

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/remote"
	"github.com/medvesek/infrastructure/lkw/src/baseconfig"
	"github.com/medvesek/infrastructure/lkw/src/constants"
)

func New() *remote.RemoteClient {
	user := baseconfig.SshUser()
	ip := baseconfig.ServerIp()

	return remote.New(user, ip)
}

func GetDestination(folder string) string {
	return fmt.Sprintf("~/%s/%s", constants.RemoteDir, folder)
}

package remote

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/utils"
)

type RemoteClient struct {
	user string
	ip   string
}

func NewRemoteClient(user string, ip string) *RemoteClient {
	return &RemoteClient{
		user,
		ip,
	}
}

func (r *RemoteClient) Cmd(cmdString string) error {
	return utils.Cmd("ssh", r.getTarget(), cmdString)
}

func (r *RemoteClient) Rsync(source string, destination string) error {
	fullDestination := fmt.Sprintf("%s:%s", r.getTarget(), destination)
	return utils.Cmd("rsync", "-avP", source, fullDestination)
}

func (r *RemoteClient) getTarget() string {
	return fmt.Sprintf("%s@%s", r.user, r.ip)
}

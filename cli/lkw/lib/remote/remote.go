package remote

import (
	"fmt"

	"github.com/medvesek/infrastructure/lkw/lib/cmd"
)

type RemoteClient struct {
	user string
	ip   string
}

func New(user string, ip string) *RemoteClient {
	return &RemoteClient{
		user,
		ip,
	}
}

func (r *RemoteClient) Cmd(cmdString string) error {
	return cmd.Run("ssh", r.getTarget(), cmdString)
}

func (r *RemoteClient) EnsureDir(path string) error {
	return r.Cmd(fmt.Sprintf("mkdir -p %s", path))
}

func (r *RemoteClient) Rsync(source string, destination string) error {
	fullDestination := fmt.Sprintf("%s:%s", r.getTarget(), destination)
	return cmd.Rsync(source, fullDestination)
}

func (r *RemoteClient) RsyncD(source string, destination string) error {
	fullDestination := fmt.Sprintf("%s:%s", r.getTarget(), destination)
	return cmd.RsyncD(source, fullDestination)
}

func (r *RemoteClient) getTarget() string {
	return fmt.Sprintf("%s@%s", r.user, r.ip)
}

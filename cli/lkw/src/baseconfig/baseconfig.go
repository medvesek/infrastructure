package baseconfig

import (
	"errors"

	"github.com/spf13/viper"
)

func CloudflareToken() string {
	return viper.GetString("cloudflare_token")
}

func ServerIp() string {
	return viper.GetString("server_ip")
}

func SshUser() string {
	return viper.GetString("ssh_user")
}

func Validate() error {
	if CloudflareToken() == "" {
		return errors.New("Cloudflare token not set")
	}
	if ServerIp() == "" {
		return errors.New("Server ip not set")
	}
	if SshUser() == "" {
		return errors.New("Ssh user not set")
	}
	return nil
}

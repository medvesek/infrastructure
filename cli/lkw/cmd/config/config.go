package config

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure lkw",
	Long:  "Configre lkw by providing base config values",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config()
	},
}

func config() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Cloudflare token: ")
	secretInput, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return err
	}

	cloudflareToken := string(secretInput)

	fmt.Print("Server ip: ")
	scanner.Scan()
	serverIp := scanner.Text()

	fmt.Print("SSH user: ")
	scanner.Scan()
	sshUser := scanner.Text()

	viperInstance := viper.New()

	viperInstance.Set("cloudflare_token", cloudflareToken)
	viperInstance.Set("server_ip", serverIp)
	viperInstance.Set("ssh_user", sshUser)

	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	if err := os.MkdirAll(home+"/.lkw", 0755); err != nil {
		return err
	}

	configFilePath := home + "/.lkw/config.yaml"
	if err := viperInstance.WriteConfigAs(configFilePath); err != nil {
		return err
	}

	if err := os.Chmod(configFilePath, 0600); err != nil {
		return err
	}

	return nil
}

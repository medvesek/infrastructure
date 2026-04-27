package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/medvesek/infrastructure/lkw/cmd/cloudflare"
	"github.com/medvesek/infrastructure/lkw/cmd/config"
	"github.com/medvesek/infrastructure/lkw/cmd/deploy"
	"github.com/medvesek/infrastructure/lkw/cmd/fun"
	"github.com/medvesek/infrastructure/lkw/cmd/ls"
	"github.com/medvesek/infrastructure/lkw/cmd/rm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "lkw",
	Short: "CLI utility for deploying to app server",
	Long:  `CLI utility for deploying to app server`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initializeConfig(cmd)
	},
}

func init() {
	rootCmd.AddCommand(cloudflare.CloudflareCmd)
	rootCmd.AddCommand(deploy.DeployCmd)
	rootCmd.AddCommand(ls.LsCmd)
	rootCmd.AddCommand(rm.RmCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(fun.FunCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initializeConfig(_ *cobra.Command) error {
	// 1. Set up Viper to use environment variables.
	viper.SetEnvPrefix("LKW")
	// Allow for nested keys in environment variables (e.g. `MYAPP_DATABASE_HOST`)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	home, err := os.UserHomeDir()

	cobra.CheckErr(err)

	viper.AddConfigPath(home + "/.lkw")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.MergeInConfig(); err != nil {
		// It's okay if the config file doesn't exist.
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	// 2. Handle the configuration file.
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("lkw.yaml")
	}

	if err := viper.MergeInConfig(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	return nil
}

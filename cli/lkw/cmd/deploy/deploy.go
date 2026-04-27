package deploy

import (
	"errors"

	"github.com/medvesek/infrastructure/lkw/src/deploy"
	"github.com/medvesek/infrastructure/lkw/src/deploy/appdeploy"
	"github.com/medvesek/infrastructure/lkw/src/deploy/staticdeploy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy to server",
	Long:  "Deploy to server",
	RunE: func(cmd *cobra.Command, args []string) error {
		deployType := viper.GetString("type")

		switch deployType {
		case "static":
			return deploy.Run(staticdeploy.New())
		case "app":
			return deploy.Run(appdeploy.New())
		default:
			return errors.New(deployType + " is not a a valid deploy type")
		}
	},
}

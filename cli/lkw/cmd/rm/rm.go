package rm

import (
	"errors"

	"github.com/medvesek/infrastructure/lkw/src/deploy"
	"github.com/medvesek/infrastructure/lkw/src/deploy/appdeploy"
	"github.com/medvesek/infrastructure/lkw/src/deploy/staticdeploy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove app",
	Long:  "Remove app",
	RunE: func(cmd *cobra.Command, args []string) error {
		deployType := viper.GetString("type")

		switch deployType {
		case "static":
			deploy.Remove(staticdeploy.New())
		case "app":
			deploy.Remove(appdeploy.New())
		default:
			return errors.New(deployType + " is not a a valid deploy type")
		}

		return nil
	},
}

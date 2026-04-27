package fun

import (
	"github.com/spf13/cobra"
)

type nype struct {
	Port        int
	Hostname    string
	Entrypoints string
}

type rype struct {
	Service string
	Expose  []nype
}

type hype struct {
	AppName string
	Expose  []rype
}

type sype struct {
	Expose []struct {
		Service     string
		Port        int
		Entrypoints string
	}
}

var FunCmd = &cobra.Command{
	Use:   "fun",
	Short: "Have fun",
	Long:  "Have fun",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

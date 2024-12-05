package config

import "github.com/spf13/cobra"

var (
	configFile string
	Cmd        = &cobra.Command{
		Use:     "config",
		Short:   "get config info",
		Example: "go run ./main.go -c ./config.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringVarP(
		&configFile, "config", "c", "./config.yml", "configuration file")
}

func run() {

}

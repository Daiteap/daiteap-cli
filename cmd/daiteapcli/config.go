package daiteapcli

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "config",
	Aliases:       []string{},
	Short:         "Command to interact with cli configurations",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			DaiteapCliPrintHelpAndExit(cmd)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

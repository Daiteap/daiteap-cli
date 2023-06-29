package daiteapcli

import (
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "project",
	Aliases:       []string{""},
	Short:         "Command to interact with projects from current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			DaiteapCliPrintHelpAndExit(cmd)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}

package daiteapcli

import (
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "profile",
	Aliases:       []string{"prof"},
	Short:         "Command to interact with your profile data",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			DaiteapCliPrintHelpAndExit(cmd)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}

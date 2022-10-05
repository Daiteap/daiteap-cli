package daiteapcli

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "user",
    Aliases: []string{},
    Short:  "Command to interact with users",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            printHelpAndExit(cmd)
        }
        return
    },
}

func init() {
    rootCmd.AddCommand(userCmd)
}
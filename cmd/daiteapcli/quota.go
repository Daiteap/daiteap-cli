package daiteapcli

import (
	"github.com/spf13/cobra"
)

var quotaCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "quota",
    Aliases: []string{},
    Short:  "Command to interact with quotas",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            printHelpAndExit(cmd)
        }
        return
    },
}

func init() {
    rootCmd.AddCommand(quotaCmd)
}
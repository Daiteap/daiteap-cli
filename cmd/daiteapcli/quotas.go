package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

var quotasCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "quotas",
    Aliases: []string{},
    Short:  "Command to interact with quotas",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            cmd.Help()
            os.Exit(0)
        }
        return
    },
}

func init() {
    rootCmd.AddCommand(quotasCmd)
}
package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "users",
    Aliases: []string{},
    Short:  "Command to interact with users",
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
    rootCmd.AddCommand(usersCmd)
}
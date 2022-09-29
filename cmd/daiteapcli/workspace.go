package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "workspace",
    Aliases: []string{"work"},
    Short:  "Command to interact with workspaces",
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
    rootCmd.AddCommand(workspaceCmd)
}
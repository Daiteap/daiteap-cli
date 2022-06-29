package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

var environmenttemplatesCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "environmenttemplates",
    Aliases: []string{"proj"},
    Short:  "Command to interact with environment templates from current tenant",
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
    rootCmd.AddCommand(environmenttemplatesCmd)
}
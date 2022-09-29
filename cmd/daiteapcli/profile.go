package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "profile",
    Aliases: []string{"prof"},
    Short:  "Command to interact with your profile data",
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
    rootCmd.AddCommand(profileCmd)
}
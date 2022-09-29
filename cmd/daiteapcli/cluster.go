package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "cluster",
    Aliases: []string{"clus"},
    Short:  "Command to interact with clusters",
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
    rootCmd.AddCommand(clusterCmd)
}
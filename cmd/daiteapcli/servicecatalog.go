package daiteapcli

import (
	"github.com/spf13/cobra"
)

var servicecatalogCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "service-catalog",
    Aliases: []string{"serv"},
    Short:  "Command to interact with service catalog.",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            printHelpAndExit(cmd)
        }
        return
    },
}

func init() {
    rootCmd.AddCommand(servicecatalogCmd)
}
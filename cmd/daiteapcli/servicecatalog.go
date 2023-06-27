package daiteapcli

import (
	"github.com/spf13/cobra"
)

func RunServicecatalogCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		DaiteapCliPrintHelpAndExit(cmd)
	}
	return
}

var servicecatalogCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "service-catalog",
	Aliases:       []string{"sc", "catalog"},
	Short:         "Command to interact with service catalog.",
	Args:          cobra.ExactArgs(0),
	Run:           RunServicecatalogCmd,
}

func init() {
	rootCmd.AddCommand(servicecatalogCmd)
}

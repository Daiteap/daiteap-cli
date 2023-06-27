package daiteapcli

import (
	"github.com/spf13/cobra"
)

var computeCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "compute",
	Aliases:       []string{},
	Short:         "Command to interact with Compute (VMs)",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			DaiteapCliPrintHelpAndExit(cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(computeCmd)
}

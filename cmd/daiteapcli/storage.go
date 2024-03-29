package daiteapcli

import (
	"github.com/spf13/cobra"
)

func RunStorageCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		DaiteapCliPrintHelpAndExit(cmd)
	}
	return
}

var storageCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "storage",
	Aliases:       []string{"stor"},
	Short:         "Command to interact with storage buckets from current tenant",
	Args:          cobra.ExactArgs(0),
	Run:           RunStorageCmd,
}

func init() {
	rootCmd.AddCommand(storageCmd)
}

package daiteapcli

import (
	"github.com/spf13/cobra"
)

func RunUserCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		DaiteapCliPrintHelpAndExit(cmd)
	}
	return
}

var userCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "user",
	Aliases:       []string{},
	Short:         "Command to interact with users",
	Args:          cobra.ExactArgs(0),
	Run:           RunUserCmd,
}

func init() {
	rootCmd.AddCommand(userCmd)
}

package daiteapcli

import (
	"github.com/spf13/cobra"
)

func RunQuotaCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		DaiteapCliPrintHelpAndExit(cmd)
	}
	return
}

var quotaCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "quota",
	Aliases:       []string{},
	Short:         "Command to interact with quotas",
	Args:          cobra.ExactArgs(0),
	Run:           RunQuotaCmd,
}

func init() {
	rootCmd.AddCommand(quotaCmd)
}

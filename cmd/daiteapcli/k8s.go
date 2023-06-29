package daiteapcli

import (
	"github.com/spf13/cobra"
)

var k8sCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "k8s",
	Aliases:       []string{},
	Short:         "Command to interact with Kubernetes environments",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			DaiteapCliPrintHelpAndExit(cmd)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)
}

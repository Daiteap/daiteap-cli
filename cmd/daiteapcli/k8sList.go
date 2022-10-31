package daiteapcli

import (
	"github.com/spf13/cobra"
)

var k8sListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list Kubernetes clusters",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		K8sList(cmd)
	},
}

func init() {
	k8sCmd.AddCommand(k8sListCmd)
}

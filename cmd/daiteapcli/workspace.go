package daiteapcli

import (
	"github.com/spf13/cobra"
)

func RunWorkspaceCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		DaiteapCliPrintHelpAndExit(cmd)
	}
	return
}

var workspaceCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "workspace",
	Aliases:       []string{"work"},
	Short:         "Command to interact with workspaces",
	Args:          cobra.ExactArgs(0),
	Run:           RunWorkspaceCmd,
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
}

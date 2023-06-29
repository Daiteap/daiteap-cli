package daiteapcli

import (
	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunWorkspaceGetCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	method := "GET"
	endpoint := ""
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := daiteapcli.JsonMarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var workspaceGetCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get",
	Aliases:       []string{},
	Short:         "Command to get active workspace for current user",
	Args:          cobra.ExactArgs(0),
	Run:           RunWorkspaceGetCmd,
}

func init() {
	workspaceCmd.AddCommand(workspaceGetCmd)
}

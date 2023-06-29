package daiteapcli

import (
	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var RunWorkspaceDetailsCmd = func(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	workspaceID, _ := cmd.Flags().GetString("workspace")
	method := "GET"
	endpoint := "/tenants/" + workspaceID
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "false", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := daiteapcli.JsonMarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var workspaceDetailsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "details",
	Aliases:       []string{},
	Short:         "Command to get workspace's detail information",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"workspace"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: RunWorkspaceDetailsCmd,
}

func init() {
	workspaceCmd.AddCommand(workspaceDetailsCmd)

	parameters := [][]interface{}{
		[]interface{}{"workspace", "ID of the workspace.", "string"},
	}

	addParameterFlags(parameters, workspaceDetailsCmd)
}

package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunWorkspaceSelect(cmd *cobra.Command, args []string) {

	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	workspaceID, _ := cmd.Flags().GetString("workspace")
	method := "POST"
	endpoint := "/user/select-tenant"
	requestBody := "{\"selectedTenant\": \"" + workspaceID + "\"}"
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, requestBody, "false", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var workspaceSelectCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "select",
	Aliases:       []string{},
	Short:         "Command to select active workspace for current user",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"workspace"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		RunWorkspaceSelect(cmd, args)
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceSelectCmd)

	parameters := [][]interface{}{
		{"workspace", "ID of the workspace.", "string"},
	}

	addParameterFlags(parameters, workspaceSelectCmd)
}

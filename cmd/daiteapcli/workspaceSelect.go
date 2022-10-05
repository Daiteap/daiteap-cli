package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

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
		workspaceID, _ := cmd.Flags().GetString("workspace")
		method := "POST"
		endpoint := "/selectTenant"
		requestBody := "{\"selectedTenant\": \"" + workspaceID + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceSelectCmd)

	parameters := [][]interface{}{
		[]interface{}{"workspace", "ID of the workspace.", "string"},
	}

	addParameterFlags(parameters, workspaceSelectCmd)
}

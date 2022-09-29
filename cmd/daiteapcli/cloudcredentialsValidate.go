package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialsValidateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "validate",
	Aliases:       []string{},
	Short:         "Command to start task which checks if cloudcredentials are valid.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		workspaceID, _ := cmd.Flags().GetString("workspace")
		cloudcredentialID, _ := cmd.Flags().GetString("cloudcredential")
		method := "POST"
		endpoint := "/validateCredentials"
		requestBody := "{\"account_id\": " + cloudcredentialID + ", \"tenant_id\": \"" + workspaceID + "\"}"
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
	cloudcredentialsCmd.AddCommand(cloudcredentialsValidateCmd)

	parameters := [][]interface{}{
		[]interface{}{"workspace", "ID of the workspace.", "string", false},
		[]interface{}{"cloudcredential", "ID of the cloud credential.", "string", false},
	}

	addParameterFlags(parameters, cloudcredentialsValidateCmd)
}
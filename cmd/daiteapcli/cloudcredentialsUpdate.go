package daiteapcli

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialsUpdateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "update",
    Aliases: []string{},
    Short:  "Command to update cloudcredentials from current tenant",
    Args: cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		provider, _ := cmd.Flags().GetString("provider")
		label, _ := cmd.Flags().GetString("label")
		description, _ := cmd.Flags().GetString("description")
		shared, _ := cmd.Flags().GetString("shared")
		method := "POST"
		endpoint := "/updateCloudCredential/" + id
		requestBody := "{\"provider\": \"" + provider + "\", \"label\": \"" + label + "\", \"description\": \"" + description + "\", \"sharedCredentials\": " + shared + "}"
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
	cloudcredentialsCmd.AddCommand(cloudcredentialsUpdateCmd)
	
	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the cloudcredentials.", "string", false},
		[]interface{}{"provider", "Provider of the cloudcredentials.", "string", false},
		[]interface{}{"label", "Label of the project.", "string", false},
		[]interface{}{"description", "Description of the cloudcredentials.", "string", false},
		[]interface{}{"shared", "Sets cloudcredentials share status.", "bool", false},
	}

	addParameterFlags(parameters, cloudcredentialsUpdateCmd)
}
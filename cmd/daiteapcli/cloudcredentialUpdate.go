package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialUpdateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "update",
	Aliases:       []string{},
	Short:         "Command to update cloudcredential from current tenant",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cloudcredential", "provider", "label", "description", "shared"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		id, _ := cmd.Flags().GetString("id")
		provider, _ := cmd.Flags().GetString("provider")
		label, _ := cmd.Flags().GetString("label")
		description, _ := cmd.Flags().GetString("description")
		shared, _ := cmd.Flags().GetString("shared")

		if provider != "google" && provider != "aws" && provider != "azure" {
			fmt.Println("Invalid provider parameter. Valid parameter values are \"google\", \"aws\" and \"azure\"")
			return
		}

		method := "POST"
		endpoint := "/cloud-credentials/" + id
		requestBody := "{\"provider\": \"" + provider + "\", \"label\": \"" + label + "\", \"description\": \"" + description + "\", \"sharedCredentials\": " + shared + "}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, verbose)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	cloudcredentialCmd.AddCommand(cloudcredentialUpdateCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloudcredential", "string"},
		[]interface{}{"provider", "cloud provider of the cloudcredential (google, aws, azure)", "string"},
		[]interface{}{"label", "label of the project", "string"},
		[]interface{}{"description", "description of the cloudcredential", "string"},
		[]interface{}{"shared", "sets cloudcredential share status", "bool"},
	}

	addParameterFlags(parameters, cloudcredentialUpdateCmd)
}

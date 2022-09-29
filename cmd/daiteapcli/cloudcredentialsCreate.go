package daiteapcli

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialsCreateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create",
	Aliases:       []string{},
	Short:         "Command to create cloudcredentials in current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		shared, _ := cmd.Flags().GetString("shared-credentials")
		label, _ := cmd.Flags().GetString("label")
		description, _ := cmd.Flags().GetString("description")
		method := "POST"
		endpoint := "/cloud-credentials"
		requestBody := ""

		if provider == "google" {
			googleKeyPath, _ := cmd.Flags().GetString("google-key")
			filename := strings.Split(googleKeyPath, "/")[len(strings.Split(googleKeyPath, "/"))-1]
			dir := strings.Split(googleKeyPath, filename)[0]
			file := fmt.Sprintf("%s/%s", dir, filename)
			content, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println("Unable to read google key file")
				return
			}
			json := strings.ReplaceAll(string(content), "\\n", "\\\\n")
			json = strings.ReplaceAll(string(json), "\n", "\\n")
			json = strings.ReplaceAll(string(json), "\"", "\\\"")

			requestBody = "{\"provider\": \"" + provider + "\", \"sharedCredentials\": " + shared + ", \"account_params\": {\"provider\": \"" + provider + "\", \"sharedCredentials\": " + shared + ", \"label\": \"" + label + "\", \"description\": \"" + description + "\", \"google_key\": \"" + json + "\"}}"
		} else if provider == "aws" {
			keyID, _ := cmd.Flags().GetString("aws-access-key-id")
			keySecret, _ := cmd.Flags().GetString("aws-secret-access-key")
			requestBody = "{\"provider\": \"" + provider + "\", \"sharedCredentials\": " + shared + ", \"account_params\": {\"provider\": \"" + provider + "\", \"sharedCredentials\": " + shared + ", \"label\": \"" + label + "\", \"description\": \"" + description + "\", \"aws_access_key_id\": \"" + keyID + "\", \"aws_secret_access_key\": \"" + keySecret + "\"}}"
		} else if provider == "azure" {
			tenantID, _ := cmd.Flags().GetString("azure-tenant-id")
			subscriptionID, _ := cmd.Flags().GetString("azure-subscription-id")
			clientID, _ := cmd.Flags().GetString("azure-client-id")
			clientSecret, _ := cmd.Flags().GetString("azure-client-secret")
			requestBody = "{\"provider\": \"" + provider + "\", \"sharedCredentials\": " + shared + ", \"account_params\": {\"provider\": \"" + provider + "\", \"sharedCredentials\": " + shared + ", \"label\": \"" + label + "\", \"description\": \"" + description + "\", \"azure_tenant_id\": \"" + tenantID + "\", \"azure_subscription_id\": \"" + subscriptionID + "\", \"azure_client_id\": \"" + clientID + "\", \"azure_client_secret\": \"" + clientSecret + "\"}}"
		} else {
			fmt.Println("Invalid provider parameter. Valid parameter values are \"google\", \"aws\" and \"azure\"")
			return
		}

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
	cloudcredentialsCmd.AddCommand(cloudcredentialsCreateCmd)

	parameters := [][]interface{}{
		[]interface{}{"provider", "cloud provider of the cloud credentials", "string", false},
		[]interface{}{"shared-credentials", "whether cloud credentials are shared", "string", false},
		[]interface{}{"label", "label of the cloud credentials.", "string", false},
		[]interface{}{"description", "description of the cloud credentials.", "string", false},
		[]interface{}{"google-key", "path to gcp credentials json file (only needed if provider is google)", "string", true},
		[]interface{}{"aws-access-key-id", "ID of the access key (only needed if provider is aws)", "string", true},
		[]interface{}{"aws-secret-access-key", "access key's secret (only needed if provider is aws)", "string", true},
		[]interface{}{"azure-tenant-id", "ID of the tenant (only needed if provider is azure)", "string", true},
		[]interface{}{"azure-subscription-id", "ID of the subscription (only needed if provider is azure)", "string", true},
		[]interface{}{"azure-client-id", "ID of the client (only needed if provider is azure)", "string", true},
		[]interface{}{"azure-client-secret", "client's secret (only needed if provider is azure)", "string", true},
	}

	addParameterFlags(parameters, cloudcredentialsCreateCmd)
}
package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var storageCreateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create",
	Aliases:       []string{},
	Short:         "Command to create storage bucket",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		credentialID, _ := cmd.Flags().GetString("credential")
		projectID, _ := cmd.Flags().GetString("project")
		name, _ := cmd.Flags().GetString("name")
		method := "POST"
		endpoint := "/buckets"
		requestBody := ""

		if provider == "google" {
			storageClass, _ := cmd.Flags().GetString("google-storage-class")
			bucketLocation, _ := cmd.Flags().GetString("google-bucket-location")
			requestBody = "{\"provider\": \"" + provider + "\", \"credential\": \"" + credentialID + "\", \"project\": \"" + projectID + "\", \"name\": \"" + name + "\", \"storage_class\": \"" + storageClass + "\", \"bucket_location\": \"" + bucketLocation + "\"}"
		} else if provider == "aws" {
			bucketLocation, _ := cmd.Flags().GetString("aws-bucket-location")
			requestBody = "{\"provider\": \"" + provider + "\", \"credential\": \"" + credentialID + "\", \"project\": \"" + projectID + "\", \"name\": \"" + name + "\", \"bucket_location\": \"" + bucketLocation + "\"}"
		} else if provider == "azure" {
			storageAccount, _ := cmd.Flags().GetString("azure-storage-account-url")
			requestBody = "{\"provider\": \"" + provider + "\", \"credential\": \"" + credentialID + "\", \"project\": \"" + projectID + "\", \"name\": \"" + name + "\", \"storage_account_url\": \"" + storageAccount + "\"}"
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
	storageCmd.AddCommand(storageCreateCmd)

	parameters := [][]interface{}{
		[]interface{}{"provider", "cloud provider in which the bucket is to be created", "string", false},
		[]interface{}{"credential", "ID of the credentials to use", "string", false},
		[]interface{}{"project", "ID of the project", "string", false},
		[]interface{}{"name", "name of the bucket", "string", false},
		[]interface{}{"google-storage-class", "storage class of the bucket (only needed if provider is google)", "string", true},
		[]interface{}{"google-bucket-location", "location of the bucket (only needed if provider is google)", "string", true},
		[]interface{}{"aws-bucket-location", "location of the bucket (only needed if provider is aws)", "string", true},
		[]interface{}{"azure-storage-account-url", "storage account url of the bucket (only needed if provider is azure)", "string", true},
	}

	addParameterFlags(parameters, storageCreateCmd)
}
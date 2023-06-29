package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunStorageCreateCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	provider, _ := cmd.Flags().GetString("provider")
	credentialID, _ := cmd.Flags().GetString("credential")
	name, _ := cmd.Flags().GetString("name")

	projectID, _ := cmd.Flags().GetString("projectID")
	if len(projectID) == 0 {
		projectName, _ := cmd.Flags().GetString("projectName")
		projectID, _ = GetProjectID(projectName)
	}

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
	}

	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, requestBody, "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var storageCreateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create",
	Aliases:       []string{},
	Short:         "Command to create storage bucket",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"provider", "credential", "name"}
		checkForRequiredFlags(requiredFlags, cmd)

		provider, _ := cmd.Flags().GetString("provider")
		if provider == "google" {
			requiredFlags = []string{"google-storage-class", "google-bucket-location"}
			checkForRequiredFlags(requiredFlags, cmd)
		} else if provider == "aws" {
			requiredFlags = []string{"aws-bucket-location"}
			checkForRequiredFlags(requiredFlags, cmd)
		} else if provider == "azure" {
			requiredFlags = []string{"azure-storage-account-url"}
			checkForRequiredFlags(requiredFlags, cmd)
		} else {
			fmt.Println("Invalid provider parameter. Valid parameter values are \"google\", \"aws\" and \"azure\"")
			DaiteapCliPrintHelpAndExit(cmd)
		}

		projectID, _ := cmd.Flags().GetString("projectID")
		projectName, _ := cmd.Flags().GetString("projectName")
		if len(projectID) == 0 && len(projectName) == 0 {
			fmt.Println("Missing or invalid project parameter")
			DaiteapCliPrintHelpAndExit(cmd)
		}

		return nil
	},
	Run: RunStorageCreateCmd,
}

func init() {
	storageCmd.AddCommand(storageCreateCmd)

	parameters := [][]interface{}{
		{"provider", "cloud provider in which the bucket is to be created (google, aws, azure)", "string"},
		{"credential", "ID of the credentials to use", "string"},
		{"projectID", "ID of the project (only needed if projectName is not set)", "string"},
		{"projectName", "ID of the project (only needed if projectID is not set)", "string"},
		{"name", "name of the bucket", "string"},
		{"google-storage-class", "storage class of the bucket (only needed if provider is google)", "string"},
		{"google-bucket-location", "location of the bucket (only needed if provider is google)", "string"},
		{"aws-bucket-location", "location of the bucket (only needed if provider is aws)", "string"},
		{"azure-storage-account-url", "storage account url of the bucket (only needed if provider is azure)", "string"},
	}

	addParameterFlags(parameters, storageCreateCmd)
}

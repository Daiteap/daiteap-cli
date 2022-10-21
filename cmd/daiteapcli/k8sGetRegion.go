package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var k8sGetRegionCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-region",
	Aliases:       []string{},
	Short:         "Command to get valid region for Kubernetes clusters",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"provider", "cloud-credential"}
		checkForRequiredFlags(requiredFlags, cmd)

		provider, _ := cmd.Flags().GetString("provider")
		if provider != "google" && provider != "aws" && provider != "azure" {
			fmt.Println("Invalid provider parameter. Valid parameter values are \"google\", \"aws\" and \"azure\"")
			printHelpAndExit(cmd)
		}

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		provider, _ := cmd.Flags().GetString("provider")
		cloudCredential, _ := cmd.Flags().GetString("cloud-credential")

		method := "POST"
		endpoint := "/getValidRegions"
		requestBody := "{\"provider\": \"" + provider + "\", \"accountId\": " + cloudCredential + "}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sGetRegionCmd)

	parameters := [][]interface{}{
		[]interface{}{"provider", "cloud provider (google, aws, azure)", "string"},
		[]interface{}{"cloud-credential", "ID of cloud credential", "string"},
	}

	addParameterFlags(parameters, k8sGetRegionCmd)
}
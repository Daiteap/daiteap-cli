package daiteapcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var k8sCreateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create",
	Aliases:       []string{},
	Short:         "Command to start task which creates Kubernetes cluster",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"lcm"}
		checkForRequiredFlags(requiredFlags, cmd)

		lcm, _ := cmd.Flags().GetString("lcm")
		if lcm == "dlcmv2" {
			templatePath, _ := cmd.Flags().GetString("dlcmv2-template")

			if len(templatePath) == 0 {
				requiredFlags := []string{"name", "description", "size", "high-availability"}
				checkForRequiredFlags(requiredFlags, cmd)

				googleCredential, _ := cmd.Flags().GetString("google-credential")
				awsCredential, _ := cmd.Flags().GetString("aws-credential")
				azureCredential, _ := cmd.Flags().GetString("azure-credential")
				if len(googleCredential) == 0 && len(awsCredential) == 0 && len(azureCredential) == 0 {
					fmt.Println("Missing or invalid credential parameter")
					printHelpAndExit(cmd)
				}

				if len(googleCredential) > 0 {
					requiredFlags := []string{"google-region"}
					checkForRequiredFlags(requiredFlags, cmd)
				}
				if len(awsCredential) > 0 {
					requiredFlags := []string{"aws-region"}
					checkForRequiredFlags(requiredFlags, cmd)
				}
				if len(azureCredential) > 0 {
					requiredFlags := []string{"azure-region"}
					checkForRequiredFlags(requiredFlags, cmd)
				}

				projectID, _ := cmd.Flags().GetString("projectID")
				projectName, _ := cmd.Flags().GetString("projectName")
				if len(projectID) == 0 && len(projectName) == 0 {
					fmt.Println("Missing or invalid project parameter")
					printHelpAndExit(cmd)
				}
			}
		} else {
			fmt.Println("Missing or invalid lcm parameter")
			printHelpAndExit(cmd)
		}

		return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		lcm, _ := cmd.Flags().GetString("lcm")

		if lcm == "dlcmv2" {
			CreateDLCMv2(cmd)
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sCreateCmd)

	parameters := [][]interface{}{
		[]interface{}{"lcm", "type of kubernetes environment", "string"},
		[]interface{}{"dlcmv2-template", "path to DLCMv2 template json file (optional)", "string"},

		[]interface{}{"projectID", "project ID in which to add the DLCMv2 environment (only needed if projectName is not set)", "string"},
		[]interface{}{"projectName", "project name in which to add the DLCMv2 environment (only needed if projectID is not set)", "string"},
		[]interface{}{"name", "name of the DLCMv2 environment", "string"},
		[]interface{}{"description", "description of the DLCMv2 environment", "string"},
		[]interface{}{"google-credential", "ID of google cloud credentials to use for the DLCMv2 environment (only needed if google provider is used)", "string"},
		[]interface{}{"google-region", "GCP region to use for the DLCMv2 environment's resources (only needed if google provider is used)", "string"},
		[]interface{}{"aws-credential", "ID of AWS cloud credentials to use for the DLCMv2 environment (only needed if aws provider is used)", "string"},
		[]interface{}{"aws-region", "AWS region to use for the DLCMv2 environment's resources (only needed if aws provider is used)", "string"},
		[]interface{}{"azure-credential", "ID of Azure cloud credentials to use for the DLCMv2 environment (only needed if azure provider is used)", "string"},
		[]interface{}{"azure-region", "Azure region to use for the DLCMv2 environment's resources (only needed if azure provider is used)", "string"},

		[]interface{}{"size", "size of the DLCMv2 environment (S, M, L, XL)", "string"},
		[]interface{}{"high-availability", "high availability DLCMv2 environment (true, false)", "string"},
	}

	addParameterFlags(parameters, k8sCreateCmd)
}
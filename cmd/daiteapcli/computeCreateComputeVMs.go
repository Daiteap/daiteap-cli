package daiteapcli

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeCreateComputeVMsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create-compute-vms",
	Aliases:       []string{},
	Short:         "Command to start task which creates compute VMs",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		templatePath, _ := cmd.Flags().GetString("compute-template")

		if len(templatePath) == 0 {
			requiredFlags := []string{"name"}
			checkForRequiredFlags(requiredFlags, cmd)

			googleCredential, _ := cmd.Flags().GetString("google-credential")
			awsCredential, _ := cmd.Flags().GetString("aws-credential")
			azureCredential, _ := cmd.Flags().GetString("azure-credential")
			if len(googleCredential) == 0 && len(awsCredential) == 0 && len(azureCredential) == 0 {
				fmt.Println("Missing or invalid credential parameter")
				printHelpAndExit(cmd)
			}

			if len(googleCredential) > 0 {
				requiredFlags := []string{"google-region", "google-vpc-cidr", "google-machine-count", "google-zone", "google-instance-type", "google-operating-system"}
				checkForRequiredFlags(requiredFlags, cmd)
			}
			if len(awsCredential) > 0 {
				requiredFlags := []string{"aws-region", "aws-vpc-cidr", "aws-machine-count", "aws-zone", "aws-instance-type", "aws-operating-system"}
				checkForRequiredFlags(requiredFlags, cmd)
			}
			if len(azureCredential) > 0 {
				requiredFlags := []string{"azure-region", "azure-vpc-cidr", "azure-machine-count", "azure-zone", "azure-instance-type", "azure-operating-system"}
				checkForRequiredFlags(requiredFlags, cmd)
			}

			projectID, _ := cmd.Flags().GetString("projectID")
			projectName, _ := cmd.Flags().GetString("projectName")
			if len(projectID) == 0 && len(projectName) == 0 {
				fmt.Println("Missing or invalid project parameter")
				printHelpAndExit(cmd)
			}
		}

		return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		templatePath, _ := cmd.Flags().GetString("compute-template")

		requestBody := ""
		if len(templatePath) > 0 {
			filename := strings.Split(templatePath, "/")[len(strings.Split(templatePath, "/"))-1]
			dir := strings.Split(templatePath, filename)[0]
			file := fmt.Sprintf("%s/%s", dir, filename)
			content, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println("Unable to read environment template file")
				os.Exit(0)
			}
			requestBody = string(content)
		} else {
			projectID, _ := cmd.Flags().GetString("projectID")
			if len(projectID) == 0 {
				projectName, _ := cmd.Flags().GetString("projectName")
				projectID, _ = GetProjectID(projectName)
			}

			name, _ := cmd.Flags().GetString("name")
			googleCredential, _ := cmd.Flags().GetString("google-credential")
			awsCredential, _ := cmd.Flags().GetString("aws-credential")
			azureCredential, _ := cmd.Flags().GetString("azure-credential")

			body := make(map[string]interface{})
			body["projectId"] = projectID
			body["internal_dns_zone"] = "daiteap.internal"
			body["clusterName"] = name
			body["onpremiseSelected"] = false
			body["alicloudSelected"] = false
			body["iotarmSelected"] = false
			body["openstackSelected"] = false

			if len(googleCredential) > 0 {
				googleRegion, _ := cmd.Flags().GetString("google-region")
				gcpCidr, _ := cmd.Flags().GetString("google-vpc-cidr")
				gcpMachineCount, _ := cmd.Flags().GetString("google-machine-count")
				gcpZone, _ := cmd.Flags().GetString("google-zone")
				gcpInstanceType, _ := cmd.Flags().GetString("google-instance-type")
				gcpOperatingSystem, _ := cmd.Flags().GetString("google-operating-system")

				body["googleSelected"] = true
				gcpValidInstanceTypes, err := GetValidInstanceTypes("google", googleCredential, googleRegion, gcpZone)
				if err != nil {
					fmt.Println("Error getting instance types")
					os.Exit(0)
				}

				count, _ := strconv.Atoi(gcpMachineCount)
				gcpNodes := make([]interface{}, count)
				for index, _ := range gcpNodes {
					gcpNode := make(map[string]interface{})
					gcpNode["is_control_plane"] = false
					gcpNode["zone"] = gcpZone
					gcpNode["instanceType"] = gcpValidInstanceTypes[gcpInstanceType]
					gcpNode["operatingSystem"] = gcpOperatingSystem

					gcpNodes[index] = gcpNode
				}

				google := make(map[string]interface{})
				google["account"], err = strconv.Atoi(googleCredential)
				if err != nil {
					fmt.Println("Error reading google credential parameter")
					os.Exit(0)
				}
				google["region"] = googleRegion
				google["vpcCidr"] = gcpCidr
				google["nodes"] = gcpNodes

				body["google"] = google
			} else {
				body["googleSelected"] = false
			}
			if len(awsCredential) > 0 {
				awsRegion, _ := cmd.Flags().GetString("aws-region")
				awsCidr, _ := cmd.Flags().GetString("aws-vpc-cidr")
				awsMachineCount, _ := cmd.Flags().GetString("aws-machine-count")
				awsZone, _ := cmd.Flags().GetString("aws-zone")
				awsInstanceType, _ := cmd.Flags().GetString("aws-instance-type")
				awsOperatingSystem, _ := cmd.Flags().GetString("aws-operating-system")

				body["awsSelected"] = true
				awsValidInstanceTypes, err := GetValidInstanceTypes("aws", awsCredential, awsRegion, awsZone)
				if err != nil {
					fmt.Println("Error getting instance types")
					os.Exit(0)
				}

				count, _ := strconv.Atoi(awsMachineCount)
				awsNodes := make([]map[string]interface{}, count)
				for index, _ := range awsNodes {
					awsNode := make(map[string]interface{})
					awsNode["is_control_plane"] = false
					awsNode["zone"] = awsZone
					awsNode["instanceType"] = awsValidInstanceTypes[awsInstanceType]
					awsNode["operatingSystem"] = awsOperatingSystem

					awsNodes[index] = awsNode
				}

				aws := make(map[string]interface{})
				aws["account"], err = strconv.Atoi(awsCredential)
				if err != nil {
					fmt.Println("Error reading aws credential parameter")
					os.Exit(0)
				}
				aws["region"] = awsRegion
				aws["vpcCidr"] = awsCidr
				aws["nodes"] = awsNodes

				body["aws"] = aws
			} else {
				body["awsSelected"] = false
			}
			if len(azureCredential) > 0 {
				azureRegion, _ := cmd.Flags().GetString("azure-region")
				azureCidr, _ := cmd.Flags().GetString("azure-vpc-cidr")
				azureMachineCount, _ := cmd.Flags().GetString("azure-machine-count")
				azureZone, _ := cmd.Flags().GetString("azure-zone")
				azureInstanceType, _ := cmd.Flags().GetString("azure-instance-type")
				azureOperatingSystem, _ := cmd.Flags().GetString("azure-operating-system")

				body["azureSelected"] = true
				azureValidInstanceTypes, err := GetValidInstanceTypes("azure", azureCredential, azureRegion, azureZone)
				if err != nil {
					fmt.Println("Error getting instance types")
					os.Exit(0)
				}

				count, _ := strconv.Atoi(azureMachineCount)
				azureNodes := make([]map[string]interface{}, count)
				for index, _ := range azureNodes {
					azureNode := make(map[string]interface{})
					azureNode["is_control_plane"] = false
					azureNode["zone"] = azureZone
					azureNode["instanceType"] = azureValidInstanceTypes[azureInstanceType]
					azureNode["operatingSystem"] = azureOperatingSystem

					azureNodes[index] = azureNode
				}

				azure := make(map[string]interface{})
				azure["account"], err = strconv.Atoi(azureCredential)
				if err != nil {
					fmt.Println("Error reading azure credential parameter")
					os.Exit(0)
				}
				azure["region"] = azureRegion
				azure["vpcCidr"] = azureCidr
				azure["nodes"] = azureNodes

				body["azure"] = azure
			} else {
				body["azureSelected"] = false
			}

			jsonBody, _ := json.Marshal(body)
			requestBody = string(jsonBody)
		}
		
		method := "POST"
		endpoint := "/createComputeVMs"
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
	computeCmd.AddCommand(computeCreateComputeVMsCmd)

	parameters := [][]interface{}{
		[]interface{}{"compute-template", "path to compute template json file (optional)", "string"},

		[]interface{}{"projectID", "project ID in which to add the Compute (VMs) environment (only needed if projectName is not set)", "string"},
		[]interface{}{"projectName", "project name in which to add the Compute (VMs) environment (only needed if projectID is not set)", "string"},
		[]interface{}{"name", "name of the Compute (VMs) environment", "string"},

		[]interface{}{"google-credential", "ID of google cloud credentials to use for the Compute (VMs) environment (only needed if google provider is used)", "string"},
		[]interface{}{"google-region", "GCP region to use for the Compute (VMs) environment's resources (only needed if google provider is used)", "string"},
		[]interface{}{"aws-credential", "ID of AWS cloud credentials to use for the Compute (VMs) environment (only needed if aws provider is used)", "string"},
		[]interface{}{"aws-region", "AWS region to use for the Compute (VMs) environment's resources (only needed if aws provider is used)", "string"},
		[]interface{}{"azure-credential", "ID of Azure cloud credentials to use for the Compute (VMs) environment (only needed if azure provider is used)", "string"},
		[]interface{}{"azure-region", "Azure region to use for the Compute (VMs) environment's resources (only needed if azure provider is used)", "string"},

		[]interface{}{"google-vpc-cidr", "google VPC CIDR of the Compute (VMs) environment (only needed if google provider is used)", "string"},
		[]interface{}{"google-machine-count", "google machine count od the Compute (VMs) environment (only needed if google provider is used)", "string"},
		[]interface{}{"google-zone", "google cloud zone for the Compute (VMs) environment (only needed if google provider is used)", "string"},
		[]interface{}{"google-instance-type", "google instance type for the Compute (VMs) environment (S, M, L, XL) (only needed if google provider is used)", "string"},
		[]interface{}{"google-operating-system", "google operating-system for the Compute (VMs) environment (only needed if google provider is used)", "string"},

		[]interface{}{"aws-vpc-cidr", "aws VPC CIDR of the Compute (VMs) environment (only needed if aws provider is used)", "string"},
		[]interface{}{"aws-machine-count", "aws machine count od the Compute (VMs) environment (only needed if aws provider is used)", "string"},
		[]interface{}{"aws-zone", "aws cloud zone for the Compute (VMs) environment (only needed if aws provider is used)", "string"},
		[]interface{}{"aws-instance-type", "aws instance type for the Compute (VMs) environment (S, M, L, XL) (only needed if aws provider is used)", "string"},
		[]interface{}{"aws-operating-system", "aws operating-system for the Compute (VMs) environment (only needed if aws provider is used)", "string"},

		[]interface{}{"azure-vpc-cidr", "azure VPC CIDR of the Compute (VMs) environment (only needed if azure provider is used)", "string"},
		[]interface{}{"azure-machine-count", "azure machine count od the Compute (VMs) environment (only needed if azure provider is used)", "string"},
		[]interface{}{"azure-zone", "azure cloud zone for the Compute (VMs) environment (only needed if azure provider is used)", "string"},
		[]interface{}{"azure-instance-type", "azure instance type for the Compute (VMs) environment (S, M, L, XL) (only needed if azure provider is used)", "string"},
		[]interface{}{"azure-operating-system", "azure operating-system for the Compute (VMs) environment (only needed if azure provider is used)", "string"},
	}

	addParameterFlags(parameters, computeCreateComputeVMsCmd)
}
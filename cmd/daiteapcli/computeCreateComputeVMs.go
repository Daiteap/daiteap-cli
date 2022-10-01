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
	Run: func(cmd *cobra.Command, args []string) {
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
			project, _ := cmd.Flags().GetString("project")
			if len(project) == 0 {
				fmt.Println("Missing or invalid \"project\" parameter")
				os.Exit(0)
			}

			name, _ := cmd.Flags().GetString("name")
			if len(name) == 0 {
				fmt.Println("Missing or invalid \"name\" parameter")
				os.Exit(0)
			}

			googleCredential, _ := cmd.Flags().GetString("google-credential")
			awsCredential, _ := cmd.Flags().GetString("aws-credential")
			azureCredential, _ := cmd.Flags().GetString("azure-credential")

			if len(googleCredential) == 0 && len(awsCredential) == 0 && len(azureCredential) == 0 {
				fmt.Println("Missing or invalid credential parameter")
				os.Exit(0)
			}

			body := make(map[string]interface{})
			body["projectId"] = project
			body["internal_dns_zone"] = "daiteap.internal"
			body["clusterName"] = name
			body["onpremiseSelected"] = false
			body["alicloudSelected"] = false
			body["iotarmSelected"] = false
			body["openstackSelected"] = false

			if len(googleCredential) > 0 {
				googleRegion, _ := cmd.Flags().GetString("google-region")
				if len(googleRegion) == 0 {
					fmt.Println("Missing or invalid \"google-region\" parameter")
					os.Exit(0)
				}

				gcpCidr, _ := cmd.Flags().GetString("google-vpc-cidr")
				if len(gcpCidr) == 0 {
					fmt.Println("Missing or invalid \"google-vpc-cidr\" parameter")
					os.Exit(0)
				}
				gcpMachineCount, _ := cmd.Flags().GetString("google-machine-count")
				if len(gcpMachineCount) == 0 {
					fmt.Println("Missing or invalid \"google-machine-count\" parameter")
					os.Exit(0)
				}
				gcpZone, _ := cmd.Flags().GetString("google-zone")
				if len(gcpZone) == 0 {
					fmt.Println("Missing or invalid \"google-zone\" parameter")
					os.Exit(0)
				}
				gcpInstanceType, _ := cmd.Flags().GetString("google-instance-type")
				if len(gcpInstanceType) == 0 {
					fmt.Println("Missing or invalid \"google-instance-type\" parameter")
					os.Exit(0)
				}
				gcpOperatingSystem, _ := cmd.Flags().GetString("google-operating-system")
				if len(gcpOperatingSystem) == 0 {
					fmt.Println("Missing or invalid \"google-operating-system\" parameter")
					os.Exit(0)
				}

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
				google["account"] = googleCredential
				google["region"] = googleRegion
				google["vpcCidr"] = gcpCidr
				google["nodes"] = gcpNodes

				body["google"] = google
			} else {
				body["googleSelected"] = false
			}
			if len(awsCredential) > 0 {
				awsRegion, _ := cmd.Flags().GetString("aws-region")
				if len(awsRegion) == 0 {
					fmt.Println("Missing or invalid \"aws-region\" parameter")
					os.Exit(0)
				}
			
				awsCidr, _ := cmd.Flags().GetString("aws-vpc-cidr")
				if len(awsCidr) == 0 {
					fmt.Println("Missing or invalid \"aws-vpc-cidr\" parameter")
					os.Exit(0)
				}
				awsMachineCount, _ := cmd.Flags().GetString("aws-machine-count")
				if len(awsMachineCount) == 0 {
					fmt.Println("Missing or invalid \"aws-machine-count\" parameter")
					os.Exit(0)
				}
				awsZone, _ := cmd.Flags().GetString("aws-zone")
				if len(awsZone) == 0 {
					fmt.Println("Missing or invalid \"aws-zone\" parameter")
					os.Exit(0)
				}
				awsInstanceType, _ := cmd.Flags().GetString("aws-instance-type")
				if len(awsInstanceType) == 0 {
					fmt.Println("Missing or invalid \"aws-instance-type\" parameter")
					os.Exit(0)
				}
				awsOperatingSystem, _ := cmd.Flags().GetString("aws-operating-system")
				if len(awsOperatingSystem) == 0 {
					fmt.Println("Missing or invalid \"aws-operating-system\" parameter")
					os.Exit(0)
				}

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
				aws["account"] = awsCredential
				aws["region"] = awsRegion
				aws["vpcCidr"] = awsCidr
				aws["nodes"] = awsNodes

				body["aws"] = aws
			} else {
				body["awsSelected"] = false
			}
			if len(azureCredential) > 0 {
				azureRegion, _ := cmd.Flags().GetString("azure-region")
				if len(azureRegion) == 0 {
					fmt.Println("Missing or invalid \"azure-region\" parameter")
					os.Exit(0)
				}
			
				azureCidr, _ := cmd.Flags().GetString("azure-vpc-cidr")
				if len(azureCidr) == 0 {
					fmt.Println("Missing or invalid \"azure-vpc-cidr\" parameter")
					os.Exit(0)
				}
				azureMachineCount, _ := cmd.Flags().GetString("azure-machine-count")
				if len(azureMachineCount) == 0 {
					fmt.Println("Missing or invalid \"azure-machine-count\" parameter")
					os.Exit(0)
				}
				azureZone, _ := cmd.Flags().GetString("azure-zone")
				if len(azureZone) == 0 {
					fmt.Println("Missing or invalid \"azure-zone\" parameter")
					os.Exit(0)
				}
				azureInstanceType, _ := cmd.Flags().GetString("azure-instance-type")
				if len(azureInstanceType) == 0 {
					fmt.Println("Missing or invalid \"azure-instance-type\" parameter")
					os.Exit(0)
				}
				azureOperatingSystem, _ := cmd.Flags().GetString("azure-operating-system")
				if len(azureOperatingSystem) == 0 {
					fmt.Println("Missing or invalid \"azure-operating-system\" parameter")
					os.Exit(0)
				}

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
				azure["account"] = azureCredential
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
	computeCmd.AddCommand(computeCreateComputeVMsCmd)

	parameters := [][]interface{}{
		[]interface{}{"compute-template", "path to compute template json file", "string", true},

		[]interface{}{"project", "project in which to add the DLCMv2 environment", "string", true},
		[]interface{}{"name", "name of the DLCMv2 environment", "string", true},

		[]interface{}{"google-credential", "ID of google cloud credentials to use for the DLCMv2 environment", "string", true},
		[]interface{}{"google-region", "GCP region to use for the DLCMv2 environment's resources", "string", true},
		[]interface{}{"aws-credential", "ID of AWS cloud credentials to use for the DLCMv2 environment", "string", true},
		[]interface{}{"aws-region", "AWS region to use for the DLCMv2 environment's resources", "string", true},
		[]interface{}{"azure-credential", "ID of Azure cloud credentials to use for the DLCMv2 environment", "string", true},
		[]interface{}{"azure-region", "Azure region to use for the DLCMv2 environment's resources", "string", true},

		[]interface{}{"google-vpc-cidr", "google VPC CIDR of the Compute (VMs) environment", "string", true},
		[]interface{}{"google-machine-count", "google machine count od the Compute (VMs) environment", "string", true},
		[]interface{}{"google-zone", "google cloud zone for the Compute (VMs) environment", "string", true},
		[]interface{}{"google-instance-type", "google instance type for the Compute (VMs) environment (S, M, L, XL)", "string", true},
		[]interface{}{"google-operating-system", "google operating-system for the Compute (VMs) environment (S, M, L, XL)", "string", true},

		[]interface{}{"aws-vpc-cidr", "aws VPC CIDR of the Compute (VMs) environment", "string", true},
		[]interface{}{"aws-machine-count", "aws machine count od the Compute (VMs) environment", "string", true},
		[]interface{}{"aws-zone", "aws cloud zone for the Compute (VMs) environment", "string", true},
		[]interface{}{"aws-instance-type", "aws instance type for the Compute (VMs) environment (S, M, L, XL)", "string", true},
		[]interface{}{"aws-operating-system", "aws operating-system for the Compute (VMs) environment (S, M, L, XL)", "string", true},

		[]interface{}{"azure-vpc-cidr", "azure VPC CIDR of the Compute (VMs) environment", "string", true},
		[]interface{}{"azure-machine-count", "azure machine count od the Compute (VMs) environment", "string", true},
		[]interface{}{"azure-zone", "azure cloud zone for the Compute (VMs) environment", "string", true},
		[]interface{}{"azure-instance-type", "azure instance type for the Compute (VMs) environment (S, M, L, XL)", "string", true},
		[]interface{}{"azure-operating-system", "azure operating-system for the Compute (VMs) environment (S, M, L, XL)", "string", true},
	}

	addParameterFlags(parameters, computeCreateComputeVMsCmd)
}
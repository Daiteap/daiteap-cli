package daiteapcli

import (
    "encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"

    "github.com/spf13/cobra"
    "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func CreateDLCMv2(cmd *cobra.Command) () {
    verbose, _ := cmd.Flags().GetString("verbose")
    templatePath, _ := cmd.Flags().GetString("dlcmv2-template")

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
        description, _ := cmd.Flags().GetString("description")
        size, _ := cmd.Flags().GetString("size")
        highAvailability, _ := cmd.Flags().GetString("high-availability")
        googleCredential, _ := cmd.Flags().GetString("google-credential")
        awsCredential, _ := cmd.Flags().GetString("aws-credential")
        azureCredential, _ := cmd.Flags().GetString("azure-credential")
        username, _ := daiteapcli.GetUsername()

        workerNodesCount := 0
        if size == "S" {
            workerNodesCount = 1
        } else if size == "M" {
            workerNodesCount = 2
        }  else if size == "L" {
            workerNodesCount = 3
        } else if size == "XL" {
            workerNodesCount = 4
        }

        controlPlaneNodesCount := 0
        if highAvailability == "false" {
            controlPlaneNodesCount = 1
        } else if highAvailability == "true" {
            controlPlaneNodesCount = 3
        }

        supportedKubernetesConfig, err := GetSupportedKubernetesConfig()
        if err != nil {
            fmt.Println("Error getting supported kubernetes configurations")
            os.Exit(0)
        }

        kubernetesConfiguration := make(map[string]interface{})
        kubernetesConfiguration["version"] = supportedKubernetesConfig["supportedKubernetesVersions"].([]interface{})[0]
        kubernetesConfiguration["serviceAddresses"] = "10.233.0.0/18"
        kubernetesConfiguration["podsSubnet"] = "10.233.64.0/18"
        kubernetesConfiguration["networkPlugin"] = supportedKubernetesConfig["supportedKubernetesNetworkPlugins"].([]interface{})[0]

        body := make(map[string]interface{})
        body["projectId"] = projectID
        body["internal_dns_zone"] = "daiteap.internal"
        body["clusterName"] = name
        body["clusterDescription"] = description
        body["onpremiseSelected"] = false
        body["alicloudSelected"] = false
        body["iotarmSelected"] = false
        body["openstackSelected"] = false
        body["type"] = 7
        body["resize"] = false
        body["kubernetesConfiguration"] = kubernetesConfiguration
        body["load_balancer_integration"] = ""

        if len(googleCredential) > 0 {
            googleRegion, _ := cmd.Flags().GetString("google-region")

            if len(body["load_balancer_integration"].(string)) == 0 {
                body["load_balancer_integration"] = "google"
            }

            body["googleSelected"] = true
            gcpValidZones, err := GetValidZones("google", googleCredential, googleRegion)
            if err != nil {
                fmt.Println("Error getting zones")
                os.Exit(0)
            }
            gcpZone := gcpValidZones[0]
            gcpValidInstanceTypes, err := GetValidInstanceTypes("google", googleCredential, googleRegion, gcpZone)
            if err != nil {
                fmt.Println("Error getting instance types")
                os.Exit(0)
            }
            gcpValidOperatingSystems, err := GetValidOperatingSystems("google", googleCredential, googleRegion, "7", username)
            if err != nil {
                fmt.Println("Error getting operating systems")
                os.Exit(0)
            }
            gcpOperatingSystem := gcpValidOperatingSystems[0]

            gcpNodes := make([]interface{}, workerNodesCount + controlPlaneNodesCount)
            for index, _ := range gcpNodes {
                gcpNode := make(map[string]interface{})
                if controlPlaneNodesCount == 0 {
                    gcpNode["is_control_plane"] = false
                } else {
                    gcpNode["is_control_plane"] = true
                    controlPlaneNodesCount -= 1
                }
                gcpNode["zone"] = gcpZone
                gcpNode["instanceType"] = gcpValidInstanceTypes[size]
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
            google["vpcCidr"] = "10.30.0.0/16"
            google["nodes"] = gcpNodes

            body["google"] = google
        } else {
            body["googleSelected"] = false
        }
        if len(awsCredential) > 0 {
            awsRegion, _ := cmd.Flags().GetString("aws-region")

            if len(body["load_balancer_integration"].(string)) == 0 {
                body["load_balancer_integration"] = "aws"
            }

            body["awsSelected"] = true
            awsValidZones, err := GetValidZones("aws", awsCredential, awsRegion)
            if err != nil {
                fmt.Println("Error getting zones")
                os.Exit(0)
            }
            awsZone := awsValidZones[0]
            awsValidInstanceTypes, err := GetValidInstanceTypes("aws", awsCredential, awsRegion, awsZone)
            if err != nil {
                fmt.Println("Error getting instance types")
                os.Exit(0)
            }
            awsValidOperatingSystems, err := GetValidOperatingSystems("aws", awsCredential, awsRegion, "7", username)
            if err != nil {
                fmt.Println("Error getting operating systems")
                os.Exit(0)
            }
            awsOperatingSystem := awsValidOperatingSystems[0]

            awsNodes := make([]interface{}, workerNodesCount + controlPlaneNodesCount)
            for index, _ := range awsNodes {
                awsNode := make(map[string]interface{})
                if controlPlaneNodesCount == 0 {
                    awsNode["is_control_plane"] = false
                } else {
                    awsNode["is_control_plane"] = true
                    controlPlaneNodesCount -= 1
                }
                awsNode["zone"] = awsZone
                awsNode["instanceType"] = awsValidInstanceTypes[size]
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
            aws["vpcCidr"] = "10.10.0.0/16"
            aws["nodes"] = awsNodes

            body["aws"] = aws
        } else {
            body["awsSelected"] = false
        }
        if len(azureCredential) > 0 {
            azureRegion, _ := cmd.Flags().GetString("azure-region")

            if len(body["load_balancer_integration"].(string)) == 0 {
                body["load_balancer_integration"] = "azure"
            }

            body["azureSelected"] = true
            azureValidZones, err := GetValidZones("azure", azureCredential, azureRegion)
            if err != nil {
                fmt.Println("Error getting zones")
                os.Exit(0)
            }
            azureZone := azureValidZones[0]
            azureValidInstanceTypes, err := GetValidInstanceTypes("azure", azureCredential, azureRegion, azureZone)
            if err != nil {
                fmt.Println("Error getting instance types")
                os.Exit(0)
            }
            azureValidOperatingSystems, err := GetValidOperatingSystems("azure", azureCredential, azureRegion, "7", username)
            if err != nil {
                fmt.Println("Error getting operating systems")
                os.Exit(0)
            }
            azureOperatingSystem := azureValidOperatingSystems[0]

            azureNodes := make([]interface{}, workerNodesCount + controlPlaneNodesCount)
            for index, _ := range azureNodes {
                azureNode := make(map[string]interface{})
                if controlPlaneNodesCount == 0 {
                    azureNode["is_control_plane"] = false
                } else {
                    azureNode["is_control_plane"] = true
                    controlPlaneNodesCount -= 1
                }
                azureNode["zone"] = azureZone
                azureNode["instanceType"] = azureValidInstanceTypes[size]
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
            azure["vpcCidr"] = "10.20.0.0/16"
            azure["nodes"] = azureNodes

            body["azure"] = azure
        } else {
            body["azureSelected"] = false
        }

        jsonBody, _ := json.Marshal(body)
        requestBody = string(jsonBody)
    
    }

    method := "POST"
    endpoint := "/createDlcmV2"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, verbose)

    if err != nil {
        fmt.Println(err)
    } else {
        output, _ := json.MarshalIndent(responseBody, "", "    ")
        fmt.Println(string(output))
    }
}
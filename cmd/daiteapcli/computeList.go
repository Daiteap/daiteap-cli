package daiteapcli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var computeListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list Compute (VMs)",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "POST"
		endpoint := "/getClusterList"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			clusterArray := make(map[string][]interface{})
			for _, cluster := range responseBody["data"].([]interface{}) {
				clusterObject := cluster.(map[string]interface{})
				switch clusterObject["type"].(float64) {
				case
					2,
					6:
					clusterArray["data"] = append(clusterArray["data"], clusterObject)
				}
			}

			if outputFormat == "json" {
				output, _ := json.MarshalIndent(clusterArray, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "Name", "Project", "Provider", "Created at", "Created by", "Status")
	
				for _, cluster := range clusterArray["data"] {
					clusterObject := cluster.(map[string]interface{})

					providers := strings.ReplaceAll(clusterObject["providers"].(string), "[", "")
					providers = strings.ReplaceAll(providers, "]", "")
					providers = strings.ReplaceAll(providers, "\"", "")
					providersArray := strings.Split(providers, ",")
					providers = ""
					for _, provider := range providersArray {
						if len(providers) == 0 {
							providers += provider
						} else {
							providers += ", " + provider
						}
					}

					installstep := clusterObject["installstep"].(float64)
					status := ""
					if installstep < 0 && installstep > -100 {
						status = "Error creating"
					} else if installstep > 0 && installstep < 100 {
						status = "Creating"
					} else if installstep == -100 {
						status = "Error deleting"
					} else if installstep == 100 {
						status = "Deleting"
					} else {
						if clusterObject["status"].(float64) == 0 {
							status = "Running"
						} else if clusterObject["status"].(float64) == 1 {
							status = "Starting"
						} else if clusterObject["status"].(float64) == -1 {
							status = "Error starting"
						} else if clusterObject["status"].(float64) == 2 {
							status = "Stopping"
						} else if clusterObject["status"].(float64) == -2 {
							status = "Error stopping"
						} else if clusterObject["status"].(float64) == 3 {
							status = "Restarting"
						} else if clusterObject["status"].(float64) == -3 {
							status = "Error restarting"
						} else if clusterObject["status"].(float64) == 10 {
							status = "Stopped"
						}
					}

					tbl.AddRow(clusterObject["id"], clusterObject["name"], clusterObject["project_name"], providers,  clusterObject["created_at"], clusterObject["contact"], status)
				}

				tbl.Print()
			} else {
				tbl := table.New("Name", "Project", "Provider", "Created at", "Created by", "Status")
	
				for _, cluster := range clusterArray["data"] {
					clusterObject := cluster.(map[string]interface{})

					providers := strings.ReplaceAll(clusterObject["providers"].(string), "[", "")
					providers = strings.ReplaceAll(providers, "]", "")
					providers = strings.ReplaceAll(providers, "\"", "")
					providersArray := strings.Split(providers, ",")
					providers = ""
					for _, provider := range providersArray {
						if len(providers) == 0 {
							providers += provider
						} else {
							providers += ", " + provider
						}
					}

					installstep := clusterObject["installstep"].(float64)
					status := ""
					if installstep < 0 && installstep > -100 {
						status = "Error creating"
					} else if installstep > 0 && installstep < 100 {
						status = "Creating"
					} else if installstep == -100 {
						status = "Error deleting"
					} else if installstep == 100 {
						status = "Deleting"
					} else {
						if clusterObject["status"].(float64) == 0 {
							status = "Running"
						} else if clusterObject["status"].(float64) == 1 {
							status = "Starting"
						} else if clusterObject["status"].(float64) == -1 {
							status = "Error starting"
						} else if clusterObject["status"].(float64) == 2 {
							status = "Stopping"
						} else if clusterObject["status"].(float64) == -2 {
							status = "Error stopping"
						} else if clusterObject["status"].(float64) == 3 {
							status = "Restarting"
						} else if clusterObject["status"].(float64) == -3 {
							status = "Error restarting"
						} else if clusterObject["status"].(float64) == 10 {
							status = "Stopped"
						}
					}

					tbl.AddRow(clusterObject["name"], clusterObject["project_name"], providers,  clusterObject["created_at"], clusterObject["contact"], status)
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	computeCmd.AddCommand(computeListCmd)
}

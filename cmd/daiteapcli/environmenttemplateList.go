package daiteapcli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var environmenttemplateListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list environment templates from current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/environment-templates"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "Name", "Description", "Providers", "Type", "Created at", "Created by")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, template := range responseBody["environmentTemplates"].([]interface{}) {
					templateObject := template.(map[string]interface{})
					providers := strings.ReplaceAll(templateObject["providers"].(string), "[", "")
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

					environmentType := ""
					switch templateObject["type"].(float64) {
					case 1:
						environmentType = "DLCM"
					case 3:
						environmentType = "DK3S"
					case 5:
						environmentType = "CAPI"
					case 7:
						environmentType = "DLCMv2"
					default:
						environmentType = "Compute (VM)"
					}
					tbl.AddRow(templateObject["id"], templateObject["name"], templateObject["description"], providers, environmentType, templateObject["created_at"], templateObject["contact"])
				}

				tbl.Print()
			} else {
				tbl := table.New("Name", "Description", "Providers", "Type", "Created at", "Created by")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, template := range responseBody["environmentTemplates"].([]interface{}) {
					templateObject := template.(map[string]interface{})
					providers := strings.ReplaceAll(templateObject["providers"].(string), "[", "")
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

					environmentType := ""
					switch templateObject["type"].(float64) {
					case 1:
						environmentType = "DLCM"
					case 3:
						environmentType = "DK3S"
					case 5:
						environmentType = "CAPI"
					case 7:
						environmentType = "DLCMv2"
					default:
						environmentType = "Compute (VM)"
					}
					tbl.AddRow(templateObject["name"], templateObject["description"], providers, environmentType, templateObject["created_at"], templateObject["contact"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	environmenttemplateCmd.AddCommand(environmenttemplateListCmd)
}

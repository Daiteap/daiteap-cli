package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var workspaceListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list workspaces for current user",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/user/active-tenants"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "false", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "Name", "Owner", "Email", "Phone", "Created at", "Updated at", "Active")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, workspace := range responseBody["activeTenants"].([]interface{}) {
					workspaceObject := workspace.(map[string]interface{})
					tbl.AddRow(workspaceObject["id"], workspaceObject["name"], workspaceObject["owner"], workspaceObject["email"], workspaceObject["phone"], workspaceObject["createdAt"], workspaceObject["updatedAt"], workspaceObject["selected"])
				}

				tbl.Print()
			} else {
				tbl := table.New("Name", "Owner", "Email", "Phone", "Created at", "Updated at", "Active")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, workspace := range responseBody["activeTenants"].([]interface{}) {
					workspaceObject := workspace.(map[string]interface{})
					tbl.AddRow(workspaceObject["name"], workspaceObject["owner"], workspaceObject["email"], workspaceObject["phone"], workspaceObject["createdAt"], workspaceObject["updatedAt"], workspaceObject["selected"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceListCmd)
}

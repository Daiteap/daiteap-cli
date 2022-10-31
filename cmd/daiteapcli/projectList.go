package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list projects from current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/projects"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "Name", "Description", "Created at", "Created by")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, project := range responseBody["data"].([]interface{}) {
					projectObject := project.(map[string]interface{})
					tbl.AddRow(projectObject["id"], projectObject["name"], projectObject["description"], projectObject["created_at"], projectObject["contact"])
				}

				tbl.Print()
			} else {
				tbl := table.New("Name", "Description", "Created at", "Created by")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, project := range responseBody["data"].([]interface{}) {
					projectObject := project.(map[string]interface{})
					tbl.AddRow(projectObject["name"], projectObject["description"], projectObject["created_at"], projectObject["contact"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}

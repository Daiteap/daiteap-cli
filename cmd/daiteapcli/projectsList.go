package daiteapcli

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
	"github.com/rodaine/table"
)

var projectsListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "list",
    Aliases: []string{},
    Short:  "Command to list projects from current tenant",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/getprojects"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else {
				tbl := table.New("Name", "Created at", "Contact")
			
				for _, project := range responseBody["projects"].([]interface{}) {
					projectObject := project.(map[string]interface{})
					tbl.AddRow(projectObject["name"], projectObject["created_at"], projectObject["contact"])
				}
			
				tbl.Print()
			}
		}
    },
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
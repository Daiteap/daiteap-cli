package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var userListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list all users in the workspace",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/getuserslist"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "User", "Role", "Projects", "Phone Number")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, user := range responseBody["users_list"].([]interface{}) {
					userObject := user.(map[string]interface{})
					projects := ""
					for _, project := range userObject["projects"].([]interface {}) {
						if len(projects) == 0 {
							projects += project.(string)
						} else {
							projects += ", " + project.(string)
						}
					}

					tbl.AddRow(userObject["id"], userObject["username"], userObject["role"], projects, userObject["phone"])
				}

				tbl.Print()
			} else {
				tbl := table.New("User", "Role", "Projects", "Phone Number")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, user := range responseBody["users_list"].([]interface{}) {
					userObject := user.(map[string]interface{})
					projects := ""
					for _, project := range userObject["projects"].([]interface {}) {
						if len(projects) == 0 {
							projects += project.(string)
						} else {
							projects += ", " + project.(string)
						}
					}

					tbl.AddRow(userObject["username"], userObject["role"], projects, userObject["phone"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}

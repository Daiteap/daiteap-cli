package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var usersListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list all users in the workspace",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/getuserslist"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else {
				tbl := table.New("User", "Role", "Projects", "Phone Number")

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
	usersCmd.AddCommand(usersListCmd)
}

package daiteap

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/cli/pkg/cli"
	"github.com/spf13/cobra"
)

var projectDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "delete",
    Aliases: []string{},
    Short:  "Command to delete project from current tenant",
	Args: cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		method := "POST"
		endpoint := "/deleteproject"
		requestBody := "{\"projectId\": \"" + id + "\"}"
		responseBody, err := daiteap.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
    },
}

func init() {
	projectCmd.AddCommand(projectDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the project.", "string", false},
	}

	addParameterFlags(parameters, projectDeleteCmd)
}
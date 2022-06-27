package daiteap

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/cli/pkg/cli"
	"github.com/spf13/cobra"
)

var projectUpdateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "update",
    Aliases: []string{},
    Short:  "Command to update project from current tenant",
	Args: cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		method := "POST"
		endpoint := "/updateProject/" + id
		requestBody := "{\"name\": \"" + name + "\", \"description\": \"" + description + "\"}"
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
	projectCmd.AddCommand(projectUpdateCmd)
	
	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the project.", "string", false},
		[]interface{}{"name", "Name of the project.", "string", false},
		[]interface{}{"description", "Description of the project.", "string", false},
	}

	addParameterFlags(parameters, projectUpdateCmd)
}
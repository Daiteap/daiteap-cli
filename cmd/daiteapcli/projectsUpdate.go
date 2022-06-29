package daiteapcli

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var projectsUpdateCmd = &cobra.Command{
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
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
    },
}

func init() {
	projectsCmd.AddCommand(projectsUpdateCmd)
	
	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the project.", "string", false},
		[]interface{}{"name", "Name of the project.", "string", false},
		[]interface{}{"description", "Description of the project.", "string", false},
	}

	addParameterFlags(parameters, projectsUpdateCmd)
}
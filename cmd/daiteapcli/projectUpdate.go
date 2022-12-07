package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var projectUpdateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "update",
	Aliases:       []string{},
	Short:         "Command to update project from current tenant",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"id", "name", "description"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		method := "PUT"
		endpoint := "/projects/" + id
		requestBody := "{\"name\": \"" + name + "\", \"description\": \"" + description + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, "true", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	projectCmd.AddCommand(projectUpdateCmd)

	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the project", "string"},
		[]interface{}{"name", "name of the project", "string"},
		[]interface{}{"description", "description of the project", "string"},
	}

	addParameterFlags(parameters, projectUpdateCmd)
}

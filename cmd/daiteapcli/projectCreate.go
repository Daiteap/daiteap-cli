package daiteapcli

import (
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var projectCreateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create",
	Aliases:       []string{},
	Short:         "Command to create project at current tenant",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"name", "description"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		method := "POST"
		endpoint := "/projects"
		requestBody := "{\"name\": \"" + name + "\", \"description\": \"" + description + "\"}"
		_, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			projectID, _ := GetProjectID(name)
			fmt.Println("New project ID: " + projectID)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectCreateCmd)

	parameters := [][]interface{}{
		[]interface{}{"name", "Name of the project.", "string"},
		[]interface{}{"description", "Description of the project.", "string"},
	}

	addParameterFlags(parameters, projectCreateCmd)
}

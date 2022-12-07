package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var projectDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete project from current tenant",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"id"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		id, _ := cmd.Flags().GetString("id")
		method := "DELETE"
		endpoint := "/projects/" + id
		requestBody := "{\"projectId\": \"" + id + "\"}"
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
	projectCmd.AddCommand(projectDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the project", "string"},
	}

	addParameterFlags(parameters, projectDeleteCmd)
}

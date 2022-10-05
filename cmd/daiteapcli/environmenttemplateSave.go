package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var environmenttemplateSaveCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "save",
	Aliases:       []string{},
	Short:         "Command to create environment template from existing environment",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"name", "environment"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		environmentID, _ := cmd.Flags().GetString("environment")
		method := "POST"
		endpoint := "/environmenttemplates/save"
		requestBody := "{\"name\": \"" + name + "\", \"environmentId\": \"" + environmentID + "\"}"
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
	environmenttemplateCmd.AddCommand(environmenttemplateSaveCmd)

	parameters := [][]interface{}{
		[]interface{}{"name", "name of the environment template", "string"},
		[]interface{}{"environment", "ID of the environment from which to create the template", "string"},
	}

	addParameterFlags(parameters, environmenttemplateSaveCmd)
}
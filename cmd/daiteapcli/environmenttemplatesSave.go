package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var environmenttemplatesSaveCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "save",
	Aliases:       []string{},
	Short:         "Command to create environment template from existing environment",
	Args:          cobra.ExactArgs(0),
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
	environmenttemplatesCmd.AddCommand(environmenttemplatesSaveCmd)

	parameters := [][]interface{}{
		[]interface{}{"name", "name of the environment template", "string", false},
		[]interface{}{"environment", "ID of the environment from which to create the template", "string", false},
	}

	addParameterFlags(parameters, environmenttemplatesSaveCmd)
}
package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var environmenttemplatesDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete environment template from current workspace.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		environmenttemplateID, _ := cmd.Flags().GetString("environmenttemplate")
		method := "POST"
		endpoint := "/environmenttemplates/delete"
		requestBody := "{\"environmentTemplateId\": \"" + environmenttemplateID + "\"}"
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
	environmenttemplatesCmd.AddCommand(environmenttemplatesDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"environmenttemplate", "ID of the environment template.", "string", false},
	}

	addParameterFlags(parameters, environmenttemplatesDeleteCmd)
}

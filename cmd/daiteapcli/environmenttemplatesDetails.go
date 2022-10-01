package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var environmenttemplatesDetailsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "details",
	Aliases:       []string{},
	Short:         "Command to get environment template's detail information.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		environmenttemplateID, _ := cmd.Flags().GetString("environmenttemplate")
		method := "GET"
		endpoint := "/environmenttemplates/get/" + environmenttemplateID
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	environmenttemplatesCmd.AddCommand(environmenttemplatesDetailsCmd)

	parameters := [][]interface{}{
		[]interface{}{"environmenttemplate", "ID of the environment template.", "string", false},
	}

	addParameterFlags(parameters, environmenttemplatesDetailsCmd)
}

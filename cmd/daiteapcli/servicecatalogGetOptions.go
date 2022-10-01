package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var servicecatalogGetOptionsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-options",
	Aliases:       []string{},
	Short:         "Command to get connection info for specific installed service",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		service, _ := cmd.Flags().GetString("service")
		method := "POST"
		endpoint := "/getServiceOptions"
		requestBody := "{\"service\": \"" + service + "\"}"
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
	servicecatalogCmd.AddCommand(servicecatalogGetOptionsCmd)

	parameters := [][]interface{}{
		[]interface{}{"service", "service which options is requested", "string", false},
	}

	addParameterFlags(parameters, servicecatalogGetOptionsCmd)
}
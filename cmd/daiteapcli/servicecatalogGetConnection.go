package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var servicecatalogGetConnectionCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-connection-info",
	Aliases:       []string{},
	Short:         "Command to get connection info for specific installed service",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("namespace")
		clusterID, _ := cmd.Flags().GetString("cluster")
		method := "POST"
		endpoint := "/getServiceConnectionInfo"
		requestBody := "{\"name\": \"" + name + "\", \"namespace\": \"" + namespace + "\", \"clusterID\": \"" + clusterID + "\"}"
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
	servicecatalogCmd.AddCommand(servicecatalogGetConnectionCmd)

	parameters := [][]interface{}{
		[]interface{}{"name", "name of the installed service.", "string", false},
		[]interface{}{"namespace", "namespace of the installed service.", "string", false},
		[]interface{}{"cluster", "ID of the cluster the service is installed on.", "string", false},
	}

	addParameterFlags(parameters, servicecatalogGetConnectionCmd)
}
package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var servicecatalogDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete installed services",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"name", "namespace", "cluster"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		name, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("namespace")
		clusterID, _ := cmd.Flags().GetString("cluster")
		method := "POST"
		endpoint := "/deleteService"
		requestBody := "{\"name\": \"" + name + "\", \"namespace\": \"" + namespace + "\", \"clusterID\": \"" + clusterID + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, verbose)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	servicecatalogCmd.AddCommand(servicecatalogDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"name", "name of the installed service.", "string"},
		[]interface{}{"namespace", "namespace of the installed service.", "string"},
		[]interface{}{"cluster", "ID of the cluster the service is installed on.", "string"},
	}

	addParameterFlags(parameters, servicecatalogDeleteCmd)
}
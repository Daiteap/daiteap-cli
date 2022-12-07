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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"name", "namespace", "cluster"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		name, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("namespace")
		clusterID, _ := cmd.Flags().GetString("cluster")
		method := "GET"
		endpoint := "/clusters/" + clusterID + "/services/" + name + "/" + namespace + "/connection-info"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	servicecatalogCmd.AddCommand(servicecatalogGetConnectionCmd)

	parameters := [][]interface{}{
		[]interface{}{"name", "name of the installed service", "string"},
		[]interface{}{"namespace", "namespace of the installed service", "string"},
		[]interface{}{"cluster", "ID of the cluster the service is installed on", "string"},
	}

	addParameterFlags(parameters, servicecatalogGetConnectionCmd)
}
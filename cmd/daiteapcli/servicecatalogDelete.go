package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunServicecatalogDeleteCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	name, _ := cmd.Flags().GetString("name")
	namespace, _ := cmd.Flags().GetString("namespace")
	clusterID, _ := cmd.Flags().GetString("cluster")
	method := "DELETE"
	endpoint := "/clusters/" + clusterID + "/services/" + name + "/" + namespace
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

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
	Run: RunServicecatalogDeleteCmd,
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

package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunServicecatalogGetOptionsCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	service, _ := cmd.Flags().GetString("service")
	method := "GET"
	endpoint := "/services/" + service + "/options"
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "false", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var servicecatalogGetOptionsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-options",
	Aliases:       []string{},
	Short:         "Command to get connection info for specific installed service",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"service"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: RunServicecatalogGetOptionsCmd,
}

func init() {
	servicecatalogCmd.AddCommand(servicecatalogGetOptionsCmd)

	parameters := [][]interface{}{
		[]interface{}{"service", "service which options is requested", "string"},
	}

	addParameterFlags(parameters, servicecatalogGetOptionsCmd)
}

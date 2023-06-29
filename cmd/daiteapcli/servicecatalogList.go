package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunServicecatalogListCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	method := "GET"
	endpoint := "/services"
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "false", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var servicecatalogListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list service catalog's services",
	Args:          cobra.ExactArgs(0),
	Run:           RunServicecatalogListCmd,
}

func init() {
	servicecatalogCmd.AddCommand(servicecatalogListCmd)
}

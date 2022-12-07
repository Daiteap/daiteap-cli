package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeGetZoneCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-zone",
	Aliases:       []string{},
	Short:         "Command to get valid zone for Compute (VMs)",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cloud-credential", "region"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		cloudCredential, _ := cmd.Flags().GetString("cloud-credential")
		region, _ := cmd.Flags().GetString("region")

		method := "GET"
		endpoint := "/cloud-credentials/" + cloudCredential + "/regions/" + region + "/zones"
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
	computeCmd.AddCommand(computeGetZoneCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloud-credential", "ID of cloud credential", "string"},
		[]interface{}{"region", "cloud region", "string"},
	}

	addParameterFlags(parameters, computeGetZoneCmd)
}
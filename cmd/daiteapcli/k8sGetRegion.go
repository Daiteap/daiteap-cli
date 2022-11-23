package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var k8sGetRegionCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-region",
	Aliases:       []string{},
	Short:         "Command to get valid region for Kubernetes clusters",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cloud-credential"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		cloudCredential, _ := cmd.Flags().GetString("cloud-credential")

		method := "GET"
		endpoint := "/cloud-credentials/" + cloudCredential + "/regions"
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
	k8sCmd.AddCommand(k8sGetRegionCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloud-credential", "ID of cloud credential", "string"},
	}

	addParameterFlags(parameters, k8sGetRegionCmd)
}
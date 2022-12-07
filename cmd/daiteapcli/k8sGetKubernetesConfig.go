package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var k8sGetKubernetesConfigCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-supported-k8s-config",
	Aliases:       []string{},
	Short:         "Command to get supported kubernetes configuration",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		method := "GET"
		endpoint := "/clusters/dlcmv2-supported-configurations"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "false", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sGetKubernetesConfigCmd)
}
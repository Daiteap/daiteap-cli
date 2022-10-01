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
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		cloudCredential, _ := cmd.Flags().GetString("cloud-credential")

		method := "POST"
		endpoint := "/getValidRegions"
		requestBody := "{\"provider\": \"" + provider + "\", \"accountId\": " + cloudCredential + "}"
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
	k8sCmd.AddCommand(k8sGetRegionCmd)

	parameters := [][]interface{}{
		[]interface{}{"provider", "cloud provider", "string", false},
		[]interface{}{"cloud-credential", "ID of cloud credential", "string", false},
	}

	addParameterFlags(parameters, k8sGetRegionCmd)
}
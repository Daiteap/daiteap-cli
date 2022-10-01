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
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		cloudCredential, _ := cmd.Flags().GetString("cloud-credential")
		region, _ := cmd.Flags().GetString("region")

		method := "POST"
		endpoint := "/getValidZones"
		requestBody := "{\"provider\": \"" + provider + "\", \"accountId\": " + cloudCredential + ",\"region\": \"" + region + "\"}"
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
	computeCmd.AddCommand(computeGetZoneCmd)

	parameters := [][]interface{}{
		[]interface{}{"provider", "cloud provider", "string", false},
		[]interface{}{"cloud-credential", "ID of cloud credential", "string", false},
		[]interface{}{"region", "cloud region", "string", false},
	}

	addParameterFlags(parameters, computeGetZoneCmd)
}
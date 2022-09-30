package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeGetOsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-os",
	Aliases:       []string{},
	Short:         "Command to get valid operating systems for Compute (VMs)",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		cloudCredential, _ := cmd.Flags().GetString("cloud-credential")
		region, _ := cmd.Flags().GetString("region")
		username, _ := daiteapcli.GetUsername()

		method := "GET"
		endpoint := "/getValidOperatingSystems/" + username + "/" + provider + "/" + cloudCredential + "/7/" + region
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	computeCmd.AddCommand(computeGetOsCmd)

	parameters := [][]interface{}{
		[]interface{}{"provider", "cloud provider", "string", false},
		[]interface{}{"cloud-credential", "ID of cloud credential", "string", false},
		[]interface{}{"region", "cloud region", "string", false},
	}

	addParameterFlags(parameters, computeGetOsCmd)
}
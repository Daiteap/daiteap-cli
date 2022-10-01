package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialsDetailsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "details",
	Aliases:       []string{},
	Short:         "Command to get cloud credentials's detail information.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cloudcredentialID, _ := cmd.Flags().GetString("cloudcredential")
		method := "GET"
		endpoint := "/cloud-credentials/" + cloudcredentialID
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
	cloudcredentialsCmd.AddCommand(cloudcredentialsDetailsCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloud credential.", "string", false},
	}

	addParameterFlags(parameters, cloudcredentialsDetailsCmd)
}
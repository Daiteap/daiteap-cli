package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialDetailsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "details",
	Aliases:       []string{},
	Short:         "Command to get cloud credential's detail information.",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cloudcredential"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
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
	cloudcredentialCmd.AddCommand(cloudcredentialDetailsCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloud credential.", "string"},
	}

	addParameterFlags(parameters, cloudcredentialDetailsCmd)
}
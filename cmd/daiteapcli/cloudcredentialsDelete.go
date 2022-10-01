package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialsDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete cloudcredentials from current workspace.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cloudcredentialID, _ := cmd.Flags().GetString("cloudcredential")
		method := "DELETE"
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
	cloudcredentialsCmd.AddCommand(cloudcredentialsDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloudcredentials.", "string", false},
	}

	addParameterFlags(parameters, cloudcredentialsDeleteCmd)
}

package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete cloudcredential from current workspace.",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cloudcredential"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		cloudcredentialID, _ := cmd.Flags().GetString("cloudcredential")
		method := "DELETE"
		endpoint := "/cloud-credentials/" + cloudcredentialID
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
	cloudcredentialCmd.AddCommand(cloudcredentialDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloudcredential", "string"},
	}

	addParameterFlags(parameters, cloudcredentialDeleteCmd)
}

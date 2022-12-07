package daiteapcli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialValidateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "validate",
	Aliases:       []string{},
	Short:         "Command to start task which checks if cloudcredential is valid.",
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

		method := "POST"
		endpoint := "/cloud-credentials/" + cloudcredentialID + "/validate"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		} else if dryRun != "false" {
			method = "GET"
    		endpoint = "/task-message/" + responseBody["taskId"].(string)
			
			for i := 0; i < 20; i++ {
				responseBody, err = daiteapcli.SendDaiteapRequest(method, endpoint, "", "false", verbose, "false")
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				if responseBody["status"] != "PENDING" {
					output, _ := json.MarshalIndent(responseBody, "", "    ")
					fmt.Println(string(output))
					os.Exit(0)
				}
				time.Sleep(time.Second * 1)
			}

			fmt.Println("Error timeout waiting for validation")
			os.Exit(0)
		}
	},
}

func init() {
	cloudcredentialCmd.AddCommand(cloudcredentialValidateCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloud credential", "string"},
	}

	addParameterFlags(parameters, cloudcredentialValidateCmd)
}
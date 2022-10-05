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
		cloudcredentialID, _ := cmd.Flags().GetString("cloudcredential")
		method := "POST"
		endpoint := "/validateCredentials"

		workspace, err := GetCurrentWorkspace()
		if err != nil {
			fmt.Println("Error getting current workspace")
			os.Exit(0)
		}

		requestBody := "{\"account_id\": " + cloudcredentialID + ", \"tenant_id\": \"" + workspace["id"] + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		taskID := responseBody["taskId"]
		endpoint = "/gettaskmessage"
		requestBody = "{\"taskId\": \"" + taskID.(string) + "\"}"
		
		for i := 0; i < 20; i++ {
			responseBody, err = daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)
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
	},
}

func init() {
	cloudcredentialCmd.AddCommand(cloudcredentialValidateCmd)

	parameters := [][]interface{}{
		[]interface{}{"cloudcredential", "ID of the cloud credential", "string"},
	}

	addParameterFlags(parameters, cloudcredentialValidateCmd)
}
package daiteapcli

import (
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeGetInstallStatusCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-install-status",
	Aliases:       []string{},
	Short:         "Command to get Compute (VMs)'s creation status",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"compute"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		clusterID, _ := cmd.Flags().GetString("compute")
		isCompute, err := IsCompute(clusterID)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if isCompute == false {
			fmt.Println("Please enter valid Compute (VMs) ID")
			os.Exit(0)
		}

		method := "POST"
		endpoint := "/getInstallationStatus"
		requestBody := "{\"ID\": \"" + clusterID + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			installStep := responseBody["installStep"].(float64)

			if installStep == 0 {
				fmt.Println("Current status: Created")
			} else if installStep >= 1 && installStep <= 3 {
				fmt.Println("Current status: Creating machines")
			} else if installStep >= 4 && installStep <= 10 {
				fmt.Println("Current status: Configuring machines")
			} else if installStep <= -1 && installStep >= -3 {
				fmt.Println("Current status: Error creating machines")
			} else if installStep <= -4 && installStep >= -10 {
				fmt.Println("Current status: Error configuring machines")
			} else if installStep == 100 {
				fmt.Println("Current status: Deleting")
			} else if installStep == -100 {
				fmt.Println("Current status: Error deleting")
			}
		}
	},
}

func init() {
	computeCmd.AddCommand(computeGetInstallStatusCmd)

	parameters := [][]interface{}{
		[]interface{}{"compute", "ID of the Compute (VMs)", "string"},
	}

	addParameterFlags(parameters, computeGetInstallStatusCmd)
}
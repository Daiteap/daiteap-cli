package daiteapcli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeCancelCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "cancel",
	Aliases:       []string{},
	Short:         "Command to cancel Compute (VMs) creation",
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
		endpoint := "/cancelClusterCreation"
		requestBody := "{\"clusterID\": \"" + clusterID + "\"}"
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
	computeCmd.AddCommand(computeCancelCmd)

	parameters := [][]interface{}{
		[]interface{}{"compute", "ID of the Compute (VMs)", "string"},
	}

	addParameterFlags(parameters, computeCancelCmd)
}
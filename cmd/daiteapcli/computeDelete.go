package daiteapcli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to start task which deletes Compute (VMs)",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"compute"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
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

		method := "DELETE"
		endpoint := "/clusters/" + clusterID + "/delete"
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
	computeCmd.AddCommand(computeDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"compute", "ID of the Compute (VMs)", "string"},
	}

	addParameterFlags(parameters, computeDeleteCmd)
}
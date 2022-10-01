package daiteapcli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var computeRenameCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "rename",
	Aliases:       []string{},
	Short:         "Command to rename Compute (VMs)",
	Args:          cobra.ExactArgs(0),
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

		name, _ := cmd.Flags().GetString("name")
		method := "POST"
		endpoint := "/renameCluster"
		requestBody := "{\"clusterID\": \"" + clusterID + "\", \"clusterName\": \"" + name + "\"}"
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
	computeCmd.AddCommand(computeRenameCmd)

	parameters := [][]interface{}{
		[]interface{}{"compute", "ID of the Compute (VMs)", "string", false},
		[]interface{}{"name", "new name of the Compute (VMs)", "string", false},
	}

	addParameterFlags(parameters, computeRenameCmd)
}
package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var clusterRenameCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "rename",
	Aliases:       []string{},
	Short:         "Command to rename cluster.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		clusterID, _ := cmd.Flags().GetString("cluster")
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
	clusterCmd.AddCommand(clusterRenameCmd)

	parameters := [][]interface{}{
		[]interface{}{"cluster", "ID of the cluster.", "string", false},
		[]interface{}{"name", "new name of the cluster.", "string", false},
	}

	addParameterFlags(parameters, clusterRenameCmd)
}
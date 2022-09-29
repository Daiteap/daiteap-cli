package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var clusterStopCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "stop",
	Aliases:       []string{},
	Short:         "Command to stop cluster",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		clusterID, _ := cmd.Flags().GetString("cluster")
		method := "POST"
		endpoint := "/stopCluster"
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
	clusterCmd.AddCommand(clusterStopCmd)

	parameters := [][]interface{}{
		[]interface{}{"cluster", "ID of the cluster.", "string", false},
	}

	addParameterFlags(parameters, clusterStopCmd)
}
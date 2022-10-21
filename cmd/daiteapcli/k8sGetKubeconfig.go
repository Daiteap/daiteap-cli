package daiteapcli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var k8sGetKubeconfigCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-kubeconfig",
	Aliases:       []string{},
	Short:         "Command to get Kubernetes cluster's kubeconfig",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cluster"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		clusterID, _ := cmd.Flags().GetString("cluster")
		isKubernetes, err := IsKubernetes(clusterID)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if isKubernetes == false {
			fmt.Println("Please enter valid Kubernetes cluster ID")
			os.Exit(0)
		}

		method := "POST"
		endpoint := "/getClusterKubeconfig"
		requestBody := "{\"clusterID\": \"" + clusterID + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sGetKubeconfigCmd)

	parameters := [][]interface{}{
		[]interface{}{"cluster", "ID of the Kubernetes cluster", "string"},
	}

	addParameterFlags(parameters, k8sGetKubeconfigCmd)
}
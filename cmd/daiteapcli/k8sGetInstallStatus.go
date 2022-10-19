package daiteapcli

import (
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var k8sGetInstallStatusCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-install-status",
	Aliases:       []string{},
	Short:         "Command to get Kubernetes cluster's creation status",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cluster"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
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
		endpoint := "/getInstallationStatus"
		requestBody := "{\"ID\": \"" + clusterID + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, verbose)

		if err != nil {
			fmt.Println(err)
		} else {
			installStep := responseBody["installStep"].(float64)

			if installStep == 0 {
				fmt.Println("Current status: Created")
			} else if installStep >= 1 && installStep <= 6 {
				fmt.Println("Current status: Allocating resources")
			} else if installStep >= 7 && installStep <= 8 {
				fmt.Println("Current status: Preparing for Kubernetes installation")
			} else if installStep >= 9 && installStep <= 12 {
				fmt.Println("Current status: Configuring machines")
			} else if installStep >= 13 && installStep <= 24 {
				fmt.Println("Current status: Installing Kubernetes")
			} else if installStep >= 25 && installStep <= 30 {
				fmt.Println("Current status: Finishing installation")
			} else if installStep <= -1 && installStep >= -6 {
				fmt.Println("Current status: Error in allocating resources")
			} else if installStep <= -7 && installStep >= -8 {
				fmt.Println("Current status: Error preparing for Kubernetes installation")
			} else if installStep <= -9 && installStep >= -12 {
				fmt.Println("Current status: Error configuring machines")
			} else if installStep <= -13 && installStep >= -24 {
				fmt.Println("Current status: Error installing Kubernetes")
			} else if installStep <= -25 && installStep >= -30 {
				fmt.Println("Current status: Error finishing installation")
			} else if installStep == 100 {
				fmt.Println("Current status: Deleting")
			} else if installStep == -100 {
				fmt.Println("Current status: Error deleting")
			}
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sGetInstallStatusCmd)

	parameters := [][]interface{}{
		[]interface{}{"cluster", "ID of the Kubernetes cluster", "string"},
	}

	addParameterFlags(parameters, k8sGetInstallStatusCmd)
}
package daiteapcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectGetCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get",
	Aliases:       []string{},
	Short:         "Command to get project resources",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		if len(id) == 0 && len(name) == 0 {
			fmt.Println("Missing or invalid project parameter")
			printHelpAndExit(cmd)
		}

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kubernetes Clusters")
		ListProjectK8s(cmd)
		fmt.Println("\n\nCompute")
		ListProjectCompute(cmd)
		fmt.Println("\n\nStorage")
		ListProjectStorage(cmd)
		fmt.Println("\n\nUsers")
		ListProjectUsers(cmd)
	},
}

func init() {
	projectCmd.AddCommand(projectGetCmd)

	parameters := [][]interface{}{
		[]interface{}{"id", "ID of the project (only needed if name is not set)", "string"},
		[]interface{}{"name", "Name of the project (only needed if id is not set)", "string"},
	}

	addParameterFlags(parameters, projectGetCmd)
}

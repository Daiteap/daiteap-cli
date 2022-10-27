package daiteapcli

import (
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "set",
	Aliases:       []string{},
	Short:         "Command to change configurations that the client uses",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"url"}
		checkForRequiredFlags(requiredFlags, cmd)

		singleUser, _ := cmd.Flags().GetString("single-user")
		if len(singleUser) != 0 {
			if singleUser != "false" && singleUser != "true" {
				fmt.Println("Invalid single-user parameter")
				printHelpAndExit(cmd)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serverURL, _ := cmd.Flags().GetString("url")
		singleUser, _ := cmd.Flags().GetString("single-user")
		if len(singleUser) == 0 {
			singleUser = "false"
		}

		err := daiteapcli.UpdateConfig(serverURL, singleUser)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Configuration updated")
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)

	parameters := [][]interface{}{
		[]interface{}{"url", "URL of the new server. Example - https://app.daiteap.com", "string"},
		[]interface{}{"single-user", "flag for single user mode (true, false). Default - false (optional)", "string"},
	}

	addParameterFlags(parameters, configSetCmd)
}

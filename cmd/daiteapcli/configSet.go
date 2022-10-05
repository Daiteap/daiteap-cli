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
		requiredFlags := []string{"server-url"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		serverURL, _ := cmd.Flags().GetString("server-url")

		err := daiteapcli.UpdateConfig(serverURL)

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
		[]interface{}{"server-url", "URL of the new server. Example - https://app.daiteap.com", "string"},
	}

	addParameterFlags(parameters, configSetCmd)
}
package daiteapcli

import (
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var configChangeServerCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "set",
	Aliases:       []string{},
	Short:         "Command to change server URL that the client uses",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"url"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serverURL, _ := cmd.Flags().GetString("url")

		err := daiteapcli.UpdateServerURL(serverURL)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Configuration updated")
		}
	},
}

func init() {
	configCmd.AddCommand(configChangeServerCmd)

	parameters := [][]interface{}{
		[]interface{}{"url", "URL of the new server. Example - https://app.daiteap.com", "string"},
	}

	addParameterFlags(parameters, configChangeServerCmd)
}

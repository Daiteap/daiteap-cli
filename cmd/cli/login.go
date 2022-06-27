package daiteap

import (
	"fmt"

	"github.com/Daiteap-D2C/cli/pkg/cli"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "login",
    Aliases: []string{},
    Short:  "Command to login and get required credentials",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		err := daiteap.Login()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Successfully Logged In.")
		}
    },
}

func init() {
    rootCmd.AddCommand(loginCmd)
}
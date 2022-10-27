package daiteapcli

import (
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "login",
	Aliases:       []string{},
	Short:         "Command to login and get required credentials",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := daiteapcli.Login()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

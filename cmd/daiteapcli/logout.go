package daiteapcli

import (
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "logout",
	Aliases:       []string{},
	Short:         "Command to logout",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := daiteapcli.Logout()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

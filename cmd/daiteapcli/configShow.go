package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var configShowCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "show",
	Aliases:       []string{},
	Short:         "Command to get configurations that the client uses",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := daiteapcli.GetConfig()

		if err != nil {
			fmt.Println(err)
		} else {
			configJson, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(configJson))
		}
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
}
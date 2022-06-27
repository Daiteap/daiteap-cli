package daiteap

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/daiteap/pkg/daiteap"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "list",
    Aliases: []string{},
    Short:  "Command to list projects from current tenant",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		method := "GET"
		endpoint := "/getprojects"
		responseBody, err := daiteap.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.Marshal(responseBody)
			fmt.Println(string(output))
		}
    },
}

func init() {
    projectCmd.AddCommand(projectListCmd)
}
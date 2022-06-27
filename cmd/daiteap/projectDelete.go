package daiteap

import (
	"fmt"
	"errors"
	"encoding/json"

	"github.com/Daiteap-D2C/daiteap/pkg/daiteap"
	"github.com/spf13/cobra"
)

var projectDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "delete",
    Aliases: []string{},
    Short:  "Command to delete project from current tenant",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a project id")
		}
		return nil
	},
    Run: func(cmd *cobra.Command, args []string) {
		method := "POST"
		endpoint := "/deleteproject"
		requestBody := "{\"projectId\": \"" + args[0] + "\"}"
		responseBody, err := daiteap.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.Marshal(responseBody)
			fmt.Println(string(output))
		}
    },
}

func init() {
    projectCmd.AddCommand(projectDeleteCmd)
}
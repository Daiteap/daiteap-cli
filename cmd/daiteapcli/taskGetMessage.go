package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var taskGetMessageCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "get-message",
	Aliases:       []string{},
	Short:         "Command to get task's message.",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := cmd.Flags().GetString("task")
		method := "POST"
		endpoint := "/gettaskmessage"
		requestBody := "{\"taskId\": \"" + taskID + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	taskCmd.AddCommand(taskGetMessageCmd)

	parameters := [][]interface{}{
		[]interface{}{"task", "ID of the task.", "string", false},
	}

	addParameterFlags(parameters, taskGetMessageCmd)
}
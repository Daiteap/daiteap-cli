package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunUserDeleteCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	username, _ := cmd.Flags().GetString("username")
	method := "DELETE"
	endpoint := "/users/" + username
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var userDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete user from the workspace",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"username"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: RunUserDeleteCmd,
}

func init() {
	userCmd.AddCommand(userDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"username", "username of the user", "string"},
	}

	addParameterFlags(parameters, userDeleteCmd)
}

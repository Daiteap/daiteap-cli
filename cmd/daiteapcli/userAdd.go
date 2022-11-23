package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var userAddCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "add",
	Aliases:       []string{},
	Short:         "Command to add user to the workspace",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"username","firstname","lastname","email","company","phone","sshpubkey","user-role"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		username, _ := cmd.Flags().GetString("username")
		firstname, _ := cmd.Flags().GetString("firstname")
		lastname, _ := cmd.Flags().GetString("lastname")
		email, _ := cmd.Flags().GetString("email")
		company, _ := cmd.Flags().GetString("company")
		phone, _ := cmd.Flags().GetString("phone")
		sshpubkey, _ := cmd.Flags().GetString("sshpubkey")
		userRole, _ := cmd.Flags().GetString("user-role")
		method := "POST"
		endpoint := "/users"
		requestBody := "{\"username\": \"" + username + "\", \"firstname\": \"" + firstname + "\", \"lastname\": \"" + lastname + "\", \"email\": \"" + email + "\", \"company\": \"" + company + "\", \"phone\": \"" + phone + "\", \"sshpubkey\": \"" + sshpubkey + "\", \"userRole\": \"" + userRole + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, "true", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	userCmd.AddCommand(userAddCmd)

	parameters := [][]interface{}{
		[]interface{}{"username", "username of the user", "string"},
		[]interface{}{"firstname", "first name of the user", "string"},
		[]interface{}{"lastname", "last name of the user", "string"},
		[]interface{}{"email", "email of the user", "string"},
		[]interface{}{"company", "company of the user", "string"},
		[]interface{}{"phone", "phone of the user", "string"},
		[]interface{}{"sshpubkey", "ssh public key of the user", "string"},
		[]interface{}{"user-role", "role of the user", "string"},
	}

	addParameterFlags(parameters, userAddCmd)
}
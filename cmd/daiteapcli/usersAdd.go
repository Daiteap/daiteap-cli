package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var usersAddCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "add",
	Aliases:       []string{},
	Short:         "Command to add user to the workspace",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		firstname, _ := cmd.Flags().GetString("firstname")
		lastname, _ := cmd.Flags().GetString("lastname")
		email, _ := cmd.Flags().GetString("email")
		company, _ := cmd.Flags().GetString("company")
		phone, _ := cmd.Flags().GetString("phone")
		sshpubkey, _ := cmd.Flags().GetString("sshpubkey")
		userRole, _ := cmd.Flags().GetString("user-role")
		method := "POST"
		endpoint := "/addnewuser"
		requestBody := "{\"username\": \"" + username + "\", \"firstname\": \"" + firstname + "\", \"lastname\": \"" + lastname + "\", \"email\": \"" + email + "\", \"company\": \"" + company + "\", \"phone\": \"" + phone + "\", \"sshpubkey\": \"" + sshpubkey + "\", \"userRole\": \"" + userRole + "\"}"
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
	usersCmd.AddCommand(usersAddCmd)

	parameters := [][]interface{}{
		[]interface{}{"username", "username of the user", "string", false},
		[]interface{}{"firstname", "first name of the user", "string", false},
		[]interface{}{"lastname", "last name of the user", "string", false},
		[]interface{}{"email", "email of the user", "string", false},
		[]interface{}{"company", "company of the user", "string", false},
		[]interface{}{"phone", "phone of the user", "string", false},
		[]interface{}{"sshpubkey", "ssh public key of the user", "string", false},
		[]interface{}{"user-role", "role of the user", "string", false},
	}

	addParameterFlags(parameters, usersAddCmd)
}
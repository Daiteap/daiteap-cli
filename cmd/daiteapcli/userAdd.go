package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunUserAddCmd(cmd *cobra.Command, args []string) {
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
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, requestBody, "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var userAddCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "add",
	Aliases:       []string{},
	Short:         "Command to add user to the workspace",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"username", "firstname", "lastname", "email", "company", "phone", "sshpubkey", "user-role"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: RunUserAddCmd,
}

func init() {
	userCmd.AddCommand(userAddCmd)

	parameters := [][]interface{}{
		{"username", "username of the user", "string"},
		{"firstname", "first name of the user", "string"},
		{"lastname", "last name of the user", "string"},
		{"email", "email of the user", "string"},
		{"company", "company of the user", "string"},
		{"phone", "phone of the user", "string"},
		{"sshpubkey", "ssh public key of the user", "string"},
		{"user-role", "role of the user", "string"},
	}

	addParameterFlags(parameters, userAddCmd)
}

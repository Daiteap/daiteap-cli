package daiteapcli

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var environmenttemplateCreateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create",
	Aliases:       []string{},
	Short:         "Command to create environment template in current workspace",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"environmenttemplate"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		environmenttemplatePath, _ := cmd.Flags().GetString("environmenttemplate")
		method := "POST"
		endpoint := "/environmenttemplates/create"

		filename := strings.Split(environmenttemplatePath, "/")[len(strings.Split(environmenttemplatePath, "/"))-1]
		dir := strings.Split(environmenttemplatePath, filename)[0]
		file := fmt.Sprintf("%s/%s", dir, filename)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Unable to read environment template file")
			return
		}
		requestBody := string(content)

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
	environmenttemplateCmd.AddCommand(environmenttemplateCreateCmd)

	parameters := [][]interface{}{
		[]interface{}{"environmenttemplate", "path to environment template json file", "string"},
	}

	addParameterFlags(parameters, environmenttemplateCreateCmd)
}
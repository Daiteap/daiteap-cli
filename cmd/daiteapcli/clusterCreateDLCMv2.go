package daiteapcli

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var clusterCreateDLCMv2Cmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "create-dlcmv2",
	Aliases:       []string{},
	Short:         "Command to start task which creates DLCMv2 cluster",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		templatePath, _ := cmd.Flags().GetString("dlcmv2-template")
		method := "POST"
		endpoint := "/createDlcmV2"

		filename := strings.Split(templatePath, "/")[len(strings.Split(templatePath, "/"))-1]
		dir := strings.Split(templatePath, filename)[0]
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
	clusterCmd.AddCommand(clusterCreateDLCMv2Cmd)

	parameters := [][]interface{}{
		[]interface{}{"dlcmv2-template", "path to DLCMv2 template json file", "string", false},
	}

	addParameterFlags(parameters, clusterCreateDLCMv2Cmd)
}
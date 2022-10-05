package daiteapcli

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var servicecatalogInstallCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "install",
	Aliases:       []string{},
	Short:         "Command to install service on kubernetes environment",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"service-name", "configuration-type", "cluster", "service-template"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, _ := cmd.Flags().GetString("service-name")
		configurationType, _ := cmd.Flags().GetString("configuration-type")
		clusterID, _ := cmd.Flags().GetString("cluster")
		templatePath, _ := cmd.Flags().GetString("service-template")
		method := "POST"
		endpoint := "/addService"

		filename := strings.Split(templatePath, "/")[len(strings.Split(templatePath, "/"))-1]
		dir := strings.Split(templatePath, filename)[0]
		file := fmt.Sprintf("%s/%s", dir, filename)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Unable to read environment template file")
			return
		}
		requestBody := "{\"serviceName\": \"" + serviceName + "\", \"configurationType\": \"" + configurationType + "\", \"clusterID\": \"" + clusterID + "\"," + string(content) + "}"
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
	servicecatalogCmd.AddCommand(servicecatalogInstallCmd)

	parameters := [][]interface{}{
		[]interface{}{"service-name", "name of the service you want to install.", "string"},
		[]interface{}{"configuration-type", "type of configuration to use for service install.", "string"},
		[]interface{}{"cluster", "ID of the cluster you want the service to be installed on.", "string"},
		[]interface{}{"service-template", "path to service template json file", "string"},
	}

	addParameterFlags(parameters, servicecatalogInstallCmd)
}
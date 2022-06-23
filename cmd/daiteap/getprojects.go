package daiteap

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"

	"github.com/Daiteap-D2C/daiteap/pkg/daiteap/utils"
	"github.com/spf13/cobra"
)

var getprojectsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "getprojects",
    Aliases: []string{},
    Short:  "Command to get projects from current tenant",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		token, err := utils.GetToken()

		daiteapServerURL := "http://localhost:8090/server"
		URL := fmt.Sprintf("%v/getprojects", daiteapServerURL)
		method := "GET"

		request, err := http.NewRequest(method, URL, nil)
		request.Header.Set("Authorization", token)

		var resp *http.Response
		var body []byte
		resp, err = http.DefaultClient.Do(request)
		if err == nil {
			body, err = ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				fmt.Println(string(body))
			} else {
				err = fmt.Errorf("invalid Status code (%v)", resp.StatusCode)
			}
		}
    },
}

func init() {
    rootCmd.AddCommand(getprojectsCmd)
}
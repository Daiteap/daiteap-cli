package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var storageDetailsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "details",
	Aliases:       []string{},
	Short:         "Command to get storage bucket's detail information",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		bucketID, _ := cmd.Flags().GetString("bucket")
		method := "GET"
		endpoint := "/buckets/" + bucketID
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	storageCmd.AddCommand(storageDetailsCmd)

	parameters := [][]interface{}{
		[]interface{}{"bucket", "ID of the bucket.", "string", false},
	}

	addParameterFlags(parameters, storageDetailsCmd)
}

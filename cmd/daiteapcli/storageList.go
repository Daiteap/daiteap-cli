package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var storageListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list storage buckets from current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/buckets"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "Name", "Cloud", "Project", "Credential", "Created At")

				for _, bucket := range responseBody["data"].([]interface{}) {
					bucketObject := bucket.(map[string]interface{})

					tbl.AddRow(bucketObject["id"], bucketObject["name"], bucketObject["provider"], bucketObject["project"], bucketObject["credential"], bucketObject["created_at"])
				}

				tbl.Print()
			} else {
				tbl := table.New("Name", "Cloud", "Project", "Credential", "Created At")

				for _, bucket := range responseBody["data"].([]interface{}) {
					bucketObject := bucket.(map[string]interface{})

					tbl.AddRow(bucketObject["name"], bucketObject["provider"], bucketObject["project"], bucketObject["credential"], bucketObject["created_at"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	storageCmd.AddCommand(storageListCmd)
}

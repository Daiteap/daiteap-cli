package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cloudcredentialListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list cloudcredentials from current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")
		dryRun, _ := cmd.Flags().GetString("dry-run")
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/cloud-credentials"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

		if err != nil {
			fmt.Println(err)
		} else if dryRun == "false" {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else if outputFormat == "wide" {
				tbl := table.New("ID", "Name", "Description", "Cloud", "Created at", "Created by")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, credential := range responseBody["data"].([]interface{}) {
					credentialObject := credential.(map[string]interface{})
					tbl.AddRow(credentialObject["id"], credentialObject["label"], credentialObject["description"], credentialObject["provider"], credentialObject["created_at"], credentialObject["contact"])
				}

				tbl.Print()
			} else {
				tbl := table.New("Name", "Description", "Cloud", "Created at", "Created by")
				headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgYellow).SprintfFunc()
				tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

				for _, credential := range responseBody["data"].([]interface{}) {
					credentialObject := credential.(map[string]interface{})
					tbl.AddRow(credentialObject["label"], credentialObject["description"], credentialObject["provider"], credentialObject["created_at"], credentialObject["contact"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	cloudcredentialCmd.AddCommand(cloudcredentialListCmd)
}

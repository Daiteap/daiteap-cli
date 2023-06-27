package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func RunStorageListCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	outputFormat, _ := cmd.Flags().GetString("output")
	method := "GET"
	endpoint := "/buckets"
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		if outputFormat == "json" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			daiteapcli.FmtPrintln(string(output))
		} else if outputFormat == "wide" {
			tbl := daiteapcli.TableNew("ID", "Name", "Cloud", "Project", "Credential", "Created At")
			headerFmt := daiteapcli.DaiteapCliColorNew(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := daiteapcli.DaiteapCliColorNew(color.FgYellow).SprintfFunc()
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, bucket := range responseBody["data"].([]interface{}) {
				bucketObject := bucket.(map[string]interface{})

				daiteapcli.TableAddRow(tbl, bucketObject["id"], bucketObject["name"], bucketObject["provider"], bucketObject["project"], bucketObject["credential"], bucketObject["created_at"])
			}

			daiteapcli.TablePrint(tbl)
		} else {
			tbl := daiteapcli.TableNew("Name", "Cloud", "Project", "Credential", "Created At")
			headerFmt := daiteapcli.DaiteapCliColorNew(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := daiteapcli.DaiteapCliColorNew(color.FgYellow).SprintfFunc()
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, bucket := range responseBody["data"].([]interface{}) {
				bucketObject := bucket.(map[string]interface{})

				daiteapcli.TableAddRow(tbl, bucketObject["name"], bucketObject["provider"], bucketObject["project"], bucketObject["credential"], bucketObject["created_at"])
			}

			daiteapcli.TablePrint(tbl)
		}
	}
}

var storageListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list storage buckets from current tenant",
	Args:          cobra.ExactArgs(0),
	Run:           RunStorageListCmd,
}

func init() {
	storageCmd.AddCommand(storageListCmd)
}

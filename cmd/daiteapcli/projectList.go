package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func RunProjectListCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	outputFormat, _ := cmd.Flags().GetString("output")
	method := "GET"
	endpoint := "/projects"
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		if outputFormat == "json" {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			daiteapcli.FmtPrintln(string(output))
		} else if outputFormat == "wide" {
			tbl := daiteapcli.TableNew("ID", "Name", "Description", "Created at", "Created by")
			headerFmt := daiteapcli.DaiteapCliColorNew(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := daiteapcli.DaiteapCliColorNew(color.FgYellow).SprintfFunc()
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, project := range responseBody["data"].([]interface{}) {
				projectObject := project.(map[string]interface{})
				daiteapcli.TableAddRow(tbl, projectObject["id"], projectObject["name"], projectObject["description"], projectObject["created_at"], projectObject["contact"])
			}

			daiteapcli.TablePrint(tbl)
		} else {
			tbl := daiteapcli.TableNew("Name", "Description", "Created at", "Created by")
			headerFmt := daiteapcli.DaiteapCliColorNew(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := daiteapcli.DaiteapCliColorNew(color.FgYellow).SprintfFunc()
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, project := range responseBody["data"].([]interface{}) {
				projectObject := project.(map[string]interface{})
				daiteapcli.TableAddRow(tbl, projectObject["name"], projectObject["description"], projectObject["created_at"], projectObject["contact"])
			}

			daiteapcli.TablePrint(tbl)
		}
	}
}

var projectListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list projects from current tenant",
	Args:          cobra.ExactArgs(0),
	Run:           RunProjectListCmd,
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}

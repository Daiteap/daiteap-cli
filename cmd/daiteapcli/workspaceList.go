package daiteapcli

import (
	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func RunWorkspaceListCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	outputFormat, _ := cmd.Flags().GetString("output")
	method := "GET"
	endpoint := "/user/active-tenants"
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "false", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		if outputFormat == "json" {
			output, _ := daiteapcli.JsonMarshalIndent(responseBody, "", "    ")
			daiteapcli.FmtPrintln(string(output))
		} else if outputFormat == "wide" {
			tbl := daiteapcli.TableNew("ID", "Name", "Owner", "Email", "Phone", "Created at", "Updated at", "Active")
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, workspace := range responseBody["activeTenants"].([]interface{}) {
				workspaceObject := workspace.(map[string]interface{})
				daiteapcli.TableAddRow(tbl, workspaceObject["id"], workspaceObject["name"], workspaceObject["owner"], workspaceObject["email"], workspaceObject["phone"], workspaceObject["createdAt"], workspaceObject["updatedAt"], workspaceObject["selected"])
			}

			daiteapcli.TablePrint(tbl)
		} else {
			tbl := daiteapcli.TableNew("Name", "Owner", "Email", "Phone", "Created at", "Updated at", "Active")
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, workspace := range responseBody["activeTenants"].([]interface{}) {
				workspaceObject := workspace.(map[string]interface{})
				daiteapcli.TableAddRow(tbl, workspaceObject["name"], workspaceObject["owner"], workspaceObject["email"], workspaceObject["phone"], workspaceObject["createdAt"], workspaceObject["updatedAt"], workspaceObject["selected"])
			}

			daiteapcli.TablePrint(tbl)
		}
	}
}

var workspaceListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list workspaces for current user",
	Args:          cobra.ExactArgs(0),
	Run:           RunWorkspaceListCmd,
}

func init() {
	workspaceCmd.AddCommand(workspaceListCmd)
}

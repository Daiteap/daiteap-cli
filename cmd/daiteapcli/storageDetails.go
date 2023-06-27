package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunStorageDetailsCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	bucketID, _ := cmd.Flags().GetString("bucket")
	method := "GET"
	endpoint := "/buckets/" + bucketID
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var storageDetailsCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "details",
	Aliases:       []string{},
	Short:         "Command to get storage bucket's detail information",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"bucket"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: RunStorageDetailsCmd,
}

func init() {
	storageCmd.AddCommand(storageDetailsCmd)

	parameters := [][]interface{}{
		[]interface{}{"bucket", "ID of the bucket.", "string"},
	}

	addParameterFlags(parameters, storageDetailsCmd)
}

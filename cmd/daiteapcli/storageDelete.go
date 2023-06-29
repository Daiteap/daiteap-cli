package daiteapcli

import (
	"encoding/json"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func RunStorageDeleteCmd(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetString("verbose")
	dryRun, _ := cmd.Flags().GetString("dry-run")
	bucketID, _ := cmd.Flags().GetString("bucket")
	method := "DELETE"
	endpoint := "/buckets/" + bucketID
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

	if err != nil {
		daiteapcli.FmtPrintln(err)
	} else if dryRun == "false" {
		output, _ := json.MarshalIndent(responseBody, "", "    ")
		daiteapcli.FmtPrintln(string(output))
	}
}

var storageDeleteCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "delete",
	Aliases:       []string{},
	Short:         "Command to delete storage bucket",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"bucket"}
		checkForRequiredFlags(requiredFlags, cmd)

		return nil
	},
	Run: RunStorageDeleteCmd,
}

func init() {
	storageCmd.AddCommand(storageDeleteCmd)

	parameters := [][]interface{}{
		[]interface{}{"bucket", "ID of the bucket.", "string"},
	}

	addParameterFlags(parameters, storageDeleteCmd)
}

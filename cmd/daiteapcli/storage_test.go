package daiteapcli_test

import (
	"testing"

	daiteapcmd "github.com/Daiteap/daiteapcli/cmd/daiteapcli"
	"github.com/spf13/cobra"
)

func TestRunStorageCmd_Success(t *testing.T) {
	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Mock the RunStorageCmd function
	printHelpAndExitCalls := 0
	daiteapcmd.DaiteapCliPrintHelpAndExit = func(command *cobra.Command) {
		printHelpAndExitCalls++
	}

	// Call the RunStorageCmd function
	daiteapcmd.RunStorageCmd(cmd, []string{})

	// Check that the RunStorageCmd function called the DaiteapCliPrintHelpAndExit function
	if printHelpAndExitCalls != 1 {
		t.Errorf("Expected RunStorageCmd to call the DaiteapCliPrintHelpAndExit function, but it didn't")
	}
}

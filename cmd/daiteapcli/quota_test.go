package daiteapcli_test

import (
	"testing"

	daiteapcmd "github.com/Daiteap/daiteapcli/cmd/daiteapcli"
	"github.com/spf13/cobra"
)

func TestRunQuotaCmd_Success(t *testing.T) {
	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Mock the RunQuotaCmd function
	printHelpAndExitCalls := 0
	daiteapcmd.DaiteapCliPrintHelpAndExit = func(command *cobra.Command) {
		printHelpAndExitCalls++
	}

	// Call the RunQuotaCmd function
	daiteapcmd.RunQuotaCmd(cmd, []string{})

	// Check that the RunQuotaCmd function called the DaiteapCliPrintHelpAndExit function
	if printHelpAndExitCalls != 1 {
		t.Errorf("Expected RunQuotaCmd to call the DaiteapCliPrintHelpAndExit function, but it didn't")
	}
}

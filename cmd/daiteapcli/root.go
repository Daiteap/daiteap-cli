package daiteapcli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var version = "0.1.6"

var rootCmd = &cobra.Command{
    Use:  "daiteapcli",
    Version: version,
    Short: "daiteapcli - CLI for Daiteap platform interaction",
    Long: `daiteapcli - CLI for Daiteap platform interaction
   
One can use daiteap to interact with Daiteap platform straight from the terminal`,
    PreRunE: func(cmd *cobra.Command, args []string) error {
        if len(args) == 0 {
            printHelpAndExit(cmd)
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "There was an error while executing your command: '%s'", err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.PersistentFlags().StringP("output", "o", "", "Specify output format.")

    var verboseFlag string
    rootCmd.PersistentFlags().StringVarP(&verboseFlag, "verbose", "v", "false", "Verbose mode.")
    rootCmd.PersistentFlags().Lookup("verbose").NoOptDefVal = "true"

    var dryRunFlag string
    rootCmd.PersistentFlags().StringVarP(&dryRunFlag, "dry-run", "d", "false", "No execute mode.")
    rootCmd.PersistentFlags().Lookup("dry-run").NoOptDefVal = "true"
}
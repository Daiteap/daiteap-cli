package daiteap

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var version = "0.0.2"

var rootCmd = &cobra.Command{
    Use:  "daiteap",
    Version: version,
    Short: "daiteap - a simple CLI to transform and inspect strings",
    Long: `daiteap is a super fancy CLI (kidding)
   
One can use daiteap to modify or inspect strings straight from the terminal`,
    PreRunE: func(cmd *cobra.Command, args []string) error {
        if len(args) == 0 {
            cmd.Help()
            os.Exit(0)
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
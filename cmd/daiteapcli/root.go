package daiteapcli

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var version = "0.1.1"

var rootCmd = &cobra.Command{
    Use:  "daiteapcli",
    Version: version,
    Short: "daiteapcli - CLI for Daiteap platform interaction",
    Long: `daiteapcli - CLI for Daiteap platform interaction
   
One can use daiteap to interact with Daiteap platform straight from the terminal`,
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
        fmt.Fprintf(os.Stderr, "There was an error while executing your command: '%s'", err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.PersistentFlags().String("output", "", "Specify output format.")
}
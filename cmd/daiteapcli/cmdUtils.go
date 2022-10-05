package daiteapcli

import (
	"os"

	"github.com/spf13/cobra"
)

func addParameterFlags(parameters [][]interface{}, command *cobra.Command) {
    for i := 0; i < len(parameters); i++ {
		if parameters[i][2].(string) == "string" {
			command.Flags().String(parameters[i][0].(string), "", parameters[i][1].(string))
		} else if parameters[i][2].(string) == "bool" {
			command.Flags().Bool(parameters[i][0].(string), false, parameters[i][1].(string))
		}
	}
}

func checkForRequiredFlags(requiredFlags []string, command *cobra.Command) {
	for _, flagName := range(requiredFlags) {
		flagValue, _ := command.Flags().GetString(flagName)
		if len(flagValue) == 0 {
			printHelpAndExit(command)
		}
	}

	return
}

func printHelpAndExit(command *cobra.Command) {
	command.Help()
	os.Exit(0)
}
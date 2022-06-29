package daiteapcli

import (
	"github.com/spf13/cobra"
)

func addParameterFlags(parameters [][]interface{}, command *cobra.Command) {
    for i := 0; i < len(parameters); i++ {
		if parameters[i][2].(string) == "string" {
			command.Flags().String(parameters[i][0].(string), "", parameters[i][1].(string))
		} else if parameters[i][2].(string) == "bool" {
			command.Flags().Bool(parameters[i][0].(string), false, parameters[i][1].(string))
		}
		if parameters[i][3].(bool) == false {
			command.MarkFlagRequired(parameters[i][0].(string))
		}
	}
}

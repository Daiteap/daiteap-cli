package daiteap

import (
	"log"

	"github.com/Daiteap-D2C/daiteap/pkg/daiteap/utils"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "login",
    Aliases: []string{},
    Short:  "Command to login and get required tokens",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        handleLoginCallback()
    },
}

func init() {
    rootCmd.AddCommand(loginCmd)
}

func handleLoginCallback() {
	utils.CloseApp.Add(1)
	config := utils.Config{
		KeycloakConfig: utils.KeycloakConfig{
			KeycloakURL: "http://localhost:8090/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: utils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}

	utils.StartServer(config)
	err := utils.OpenBrowser(utils.BuildAuthorizationRequest(config))
	if err != nil {
		log.Fatalf("Could not open the browser for url %v", utils.BuildAuthorizationRequest(config))
	}

	utils.CloseApp.Wait()
}
package authUtils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func BuildAuthorizationRequest(config Config) string {
	return fmt.Sprintf(
		"%v/realms/%v/protocol/openid-connect/auth?client_id=%v&redirect_uri=%v&response_mode=query&response_type=code&scope=openid",
		config.KeycloakConfig.KeycloakURL,
		config.KeycloakConfig.Realm,
		config.KeycloakConfig.ClientID,
		config.EmbeddedServerConfig.GetCallbackURL(),
	)
}

func BuildTokenExchangeRequest(config Config, code string) (*http.Request, error) {
	tokenURL := fmt.Sprintf("%v/realms/%v/protocol/openid-connect/token", config.KeycloakConfig.KeycloakURL, config.KeycloakConfig.Realm)

	body := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"client_id":    {config.KeycloakConfig.ClientID},
		"redirect_uri": {config.EmbeddedServerConfig.GetCallbackURL()},
	}

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, err
}

func BuildRefreshRequest(keycloakConfig KeycloakConfig, refreshToken string) (*http.Request, error) {
	refreshURL := fmt.Sprintf(
		"%v/realms/%v/protocol/openid-connect/token",
		keycloakConfig.KeycloakURL,
		keycloakConfig.Realm,
	)

	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)
	form.Add("client_id", keycloakConfig.ClientID)

	req, err := http.NewRequest("POST", refreshURL, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

func BuildLogoutRequest(keycloakConfig KeycloakConfig, accessToken string, refreshToken string) (*http.Request, error) {
	logoutURL := fmt.Sprintf(
		"%v/realms/%v/protocol/openid-connect/logout",
		keycloakConfig.KeycloakURL,
		keycloakConfig.Realm,
	)

	form := url.Values{}
	form.Add("client_id", keycloakConfig.ClientID)
	form.Add("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", logoutURL, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer " + accessToken)

	return req, err
}
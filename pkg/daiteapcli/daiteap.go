package daiteapcli

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli/authUtils"
)

func GetActiveToken() (string, error) {
	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: "https://app.daiteap.com/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}
	authConfig, err := authUtils.GetConfig()

	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}

	accessToken := authConfig.AccessToken
	refreshToken := authConfig.RefreshToken
	expired, err := authUtils.IsTokenExpired(&accessToken)

	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}

	if expired == true {
		expired, err = authUtils.IsTokenExpired(&refreshToken)

		if err != nil {
			err := fmt.Errorf("Error reading token. Please login again.")
			return "", err
		}

		if expired == true {
			err := fmt.Errorf("Your credentials are expired. Please login again.")
			return "", err
		}

		err = authUtils.RefreshAccessToken(&config)
		if err != nil {
			err := fmt.Errorf("Error refreshing accessToken. Please login again.")
			return "", err
		}
		authConfig, err = authUtils.GetConfig()
		accessToken = authConfig.AccessToken
	}

	return accessToken, nil
}

func Login() error {
	authUtils.CloseApp.Add(1)
	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: "https://app.daiteap.com/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}

	authUtils.StartServer(config)
	err := authUtils.OpenBrowser(authUtils.BuildAuthorizationRequest(config))
	if err != nil {
		err := fmt.Errorf("Could not open the browser for url %v", authUtils.BuildAuthorizationRequest(config))
		return err
	}

	authUtils.CloseApp.Wait()

	return nil
}

func SendDaiteapRequest(method string, endpoint string, requestBody string) (map[string]interface{}, error) {
	var resp *http.Response
	var responseBody []byte
	emptyResponseBody := make(map[string]interface{})
	daiteapServerURL := "https://app.daiteap.com/server"
	URL := fmt.Sprintf("%v"+endpoint, daiteapServerURL)

	token, err := GetActiveToken()
	if err != nil {
		return emptyResponseBody, err
	}

	request, err := http.NewRequest(method, URL, strings.NewReader(requestBody))
	request.Header.Set("Authorization", token)
	request.Header.Set("Content-type", "application/json")

	resp, err = http.DefaultClient.Do(request)
	if err == nil {
		responseBody, err = ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			var f interface{}
			json.Unmarshal(responseBody, &f)
			switch f.(type) {
			case []interface{}:
				arrayResponseBody := make(map[string]interface{})
				arrayResponseBody["data"] = f
				return arrayResponseBody, nil
			}

			m := f.(map[string]interface{})

			return m, nil
		} else {
			err = fmt.Errorf("invalid Status code (%v)", resp.StatusCode)
			return emptyResponseBody, err
		}
	} else {
		return emptyResponseBody, err
	}
}

func GetUsername() (string, error) {
	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: "https://app.daiteap.com/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}
	authConfig, err := authUtils.GetConfig()

	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}

	accessToken := authConfig.AccessToken
	refreshToken := authConfig.RefreshToken
	expired, err := authUtils.IsTokenExpired(&accessToken)

	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}

	if expired == true {
		expired, err = authUtils.IsTokenExpired(&refreshToken)

		if err != nil {
			err := fmt.Errorf("Error reading token. Please login again.")
			return "", err
		}

		if expired == true {
			err := fmt.Errorf("Your credentials are expired. Please login again.")
			return "", err
		}

		err = authUtils.RefreshAccessToken(&config)
		if err != nil {
			err := fmt.Errorf("Error refreshing accessToken. Please login again.")
			return "", err
		}
		authConfig, err = authUtils.GetConfig()
		accessToken = authConfig.AccessToken
	}

	encodedTokenPayload := strings.Split(accessToken, ".")[1]
	if len(encodedTokenPayload)%4 == 3 {
        encodedTokenPayload += "="
    } else if len(encodedTokenPayload)%4 == 2 {
        encodedTokenPayload += "=="
    } else if len(encodedTokenPayload)%4 == 1 {
        encodedTokenPayload += "==="
    }

	payload, _ := base64.StdEncoding.DecodeString(encodedTokenPayload)

	var jsonMap map[string]interface{}
	json.Unmarshal(payload, &jsonMap)
	username := jsonMap["preferred_username"].(string)

	return username, nil
}
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
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func GetActiveToken() (string, error) {
	authConfig, err := authUtils.GetConfig()
	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}
	if !(len(authConfig.ServerURL) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return "", err
	}

	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: authConfig.ServerURL + "/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}

	accessToken := authConfig.AccessToken
	refreshToken := authConfig.RefreshToken
	expired := true
	if len(accessToken) > 0 {
		expired, err = authUtils.IsTokenExpired(&accessToken)
		if err != nil {
			err := fmt.Errorf("Error reading token. Please login again.")
			return "", err
		}
	} else {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}

	if expired == true {
		if len(refreshToken) > 0 {
			expired, err = authUtils.IsTokenExpired(&refreshToken)
			if err != nil {
				err := fmt.Errorf("Error reading token. Please login again.")
				return "", err
			}
		} else {
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
	authConfig, err := authUtils.GetConfig()
	if err != nil {
		err := fmt.Errorf("Error reading config. Please update it.")
		return err
	}
	if !(len(authConfig.ServerURL) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return err
	}

	if authConfig.SingleUser == "true" {
		fmt.Println("Single user mode: No login required.")
		return nil
	}

	authUtils.CloseApp.Add(1)
	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: authConfig.ServerURL + "/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}

	authUtils.StartServer(config)
	err = authUtils.OpenBrowser(authUtils.BuildAuthorizationRequest(config))
	if err != nil {
		err := fmt.Errorf("Could not open the browser for url %v", authUtils.BuildAuthorizationRequest(config))
		return err
	}

	authUtils.CloseApp.Wait()

	fmt.Println("Successfully Logged In.")
	return nil
}

func Logout() error {
	authConfig, err := authUtils.GetConfig()
	if err != nil {
		err := fmt.Errorf("Error reading config. Please update it.")
		return err
	}
	if !(len(authConfig.ServerURL) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return err
	}

	if authConfig.SingleUser == "true" {
		fmt.Println("Single user mode: No login required.")
		return nil
	}

	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: authConfig.ServerURL + "/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}

	err = authUtils.Logout(&config)
	if err != nil {
		err := fmt.Errorf("Error logging out.")
		return err
	}

	fmt.Println("Successfully Logged Out.")
	return nil
}

var FmtPrintln = fmt.Println
var JsonMarshalIndent = json.MarshalIndent
var DaiteapCliColorNew = color.New
var DaiteapcliSendDaiteapRequest = SendDaiteapRequest

var TableNew = table.New
var TablePrint = func(t table.Table) {
	t.Print()
}
var TableAddRow = func(t table.Table, row ...interface{}) {
	t.AddRow(row...)
}

func SendDaiteapRequest(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
	var resp *http.Response
	var responseBody []byte
	emptyResponseBody := make(map[string]interface{})

	authConfig, err := authUtils.GetConfig()
	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return emptyResponseBody, err
	}
	if !(len(authConfig.ServerURL) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return emptyResponseBody, err
	}

	daiteapServerURL := authConfig.ServerURL + "/server"
	if tenant == "true" {
		workspace, err := getActiveWorkspace()
		if err != nil {
			err := fmt.Errorf("Error getting selected workspace. Please try logging in again.")
			return emptyResponseBody, err
		}
		daiteapServerURL = daiteapServerURL + "/tenants/" + workspace
	}
	URL := fmt.Sprintf("%v"+endpoint, daiteapServerURL)

	request, err := http.NewRequest(method, URL, strings.NewReader(requestBody))
	if authConfig.SingleUser == "false" {
		token, err := GetActiveToken()
		if err != nil {
			return emptyResponseBody, err
		}
		request.Header.Set("Authorization", token)
		request.Header.Set("Content-type", "application/json")
	}

	if dryRun != "false" {
		fmt.Println("URL:")
		fmt.Println(URL)
		fmt.Println("\nMethod:")
		fmt.Println(method)
		fmt.Println("\nHeaders:")
		headers, _ := json.Marshal(request.Header)
		fmt.Println(string(headers))
		fmt.Println("\nBody:")
		fmt.Println(requestBody)

		return emptyResponseBody, nil
	}

	resp, err = http.DefaultClient.Do(request)
	if err == nil {
		responseBody, err = ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
		if verbose != "false" {
			fmt.Println(string(responseBody) + "\n\n")
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode <= 300 {
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

func getActiveWorkspace() (string, error) {
	method := "GET"
	endpoint := "/user/active-tenants"
	responseBody, err := SendDaiteapRequest(method, endpoint, "", "false", "false", "false")

	if err != nil {
		return "", err
	} else {
		workspace := fmt.Sprint(responseBody["selectedTenant"])
		return workspace, nil
	}
}

func GetUsername() (string, error) {
	authConfig, err := authUtils.GetConfig()
	if err != nil {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}
	if !(len(authConfig.ServerURL) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return "", err
	}

	config := authUtils.Config{
		KeycloakConfig: authUtils.KeycloakConfig{
			KeycloakURL: authConfig.ServerURL + "/auth",
			Realm:       "Daiteap",
			ClientID:    "daiteap-cli",
		},
		EmbeddedServerConfig: authUtils.EmbeddedServerConfig{
			Port:         3000,
			CallbackPath: "sso-callback",
		},
	}

	accessToken := authConfig.AccessToken
	refreshToken := authConfig.RefreshToken
	expired := true
	if len(accessToken) > 0 {
		expired, err = authUtils.IsTokenExpired(&accessToken)
		if err != nil {
			err := fmt.Errorf("Error reading token. Please login again.")
			return "", err
		}
	} else {
		err := fmt.Errorf("Error reading token. Please login again.")
		return "", err
	}

	if expired == true {
		if len(refreshToken) > 0 {
			expired, err = authUtils.IsTokenExpired(&refreshToken)
			if expired == true {
				err := fmt.Errorf("Your credentials are expired. Please login again.")
				return "", err
			}
		} else {
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

func UpdateConfig(serverURL string, singleUser string) error {
	var cfg *authUtils.IConfig = &authUtils.IConfig{
		AccessToken:  "",
		RefreshToken: "",
		ServerURL:    serverURL,
		SingleUser:   singleUser,
	}

	err := authUtils.SaveConfig(cfg)
	if err != nil {
		err := fmt.Errorf("Error saving configuration.")
		return err
	}

	return nil
}

func GetConfig() (map[string]interface{}, error) {
	config := make(map[string]interface{})
	authConfig, err := authUtils.GetConfig()
	if err != nil {
		err := fmt.Errorf("Error reading config. Please update it.")
		return config, err
	}
	if !(len(authConfig.ServerURL) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return config, err
	}
	if !(len(authConfig.SingleUser) > 0) {
		err := fmt.Errorf("Error reading configuration. Please update it.")
		return config, err
	}

	config["Server URL"] = authConfig.ServerURL
	config["Single user mode"] = authConfig.SingleUser

	return config, nil
}

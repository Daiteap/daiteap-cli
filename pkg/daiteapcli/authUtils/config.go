package authUtils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type IConfig struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ServerURL string `json:"server_url"`
}

func getConfigLocation() (string, error) {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfgPath, "daiteap"), nil
}

func InitConfig() error {
	daiteapCfgDir, err := getConfigLocation()
	if err != nil {
		return err
	}

	// create daiteap config directory
	if _, err = os.Stat(daiteapCfgDir); os.IsNotExist(err) {
		err = os.MkdirAll(daiteapCfgDir, 0o700)
		if err != nil {
			return err
		}
	}

	return err
}

func SaveConfig(cfg *IConfig) error {
	cfgDir, err := getConfigLocation()
	if err != nil {
		return err
	}

	file := fmt.Sprintf("%s/%s", cfgDir, "config.json")
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}

	cfgParentDir := strings.Split(cfgDir, "/daiteap")[0]
	if _, err = os.Stat(cfgParentDir); os.IsNotExist(err) {
		err = os.Mkdir(cfgParentDir, 0o700)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(cfgDir); os.IsNotExist(err) {
		err = os.Mkdir(cfgDir, 0o700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0o600)

	if err != nil {
		return fmt.Errorf("%v: %w", "unable to save config", err)
	}
	return nil
}

func GetConfig() (IConfig, error) {
	cfgDir, err := getConfigLocation()
	if err != nil {
		return IConfig{}, err
	}

	file := fmt.Sprintf("%s/%s", cfgDir, "config.json")

	content, err := ioutil.ReadFile(file)
	if err != nil {
		var cfg IConfig = IConfig{
			AccessToken:  "",
			RefreshToken: "",
			ServerURL:    "https://app.daiteap.com",
		}

		err := SaveConfig(&cfg)
		if err != nil {
			err := fmt.Errorf("Error saving configuration.")
			return IConfig{}, err
		}
		content, err = ioutil.ReadFile(file)
		if err != nil {
			return IConfig{}, fmt.Errorf("%v: %w", "unable to read config", err)
		}
	}

	var f interface{}
	json.Unmarshal(content, &f)
	m := f.(map[string]interface{})

	if _, ok := m["access_token"]; !ok {
		return IConfig{}, fmt.Errorf("%v: %w", "unable to read config", err)
	}
	if _, ok := m["refresh_token"]; !ok {
		return IConfig{}, fmt.Errorf("%v: %w", "unable to read config", err)
	}
	if _, ok := m["server_url"]; !ok {
		return IConfig{}, fmt.Errorf("%v: %w", "unable to read config", err)
	}
	var cfg IConfig = IConfig{
		AccessToken:  m["access_token"].(string),
		RefreshToken: m["refresh_token"].(string),
		ServerURL:    m["server_url"].(string),
	}

	return cfg, nil
}
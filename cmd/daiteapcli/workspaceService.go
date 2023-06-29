package daiteapcli

import (
	"errors"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func GetCurrentWorkspace() (map[string]string, error) {
	method := "GET"
	endpoint := ""
	responseBody, err := daiteapcli.DaiteapcliSendDaiteapRequest(method, endpoint, "", "true", "false", "false")

	if err != nil {
		return nil, err
	}

	tenantObj, ok := responseBody["tenant"].(map[string]interface{})
	if !ok {
		return nil, errors.New("Missing or invalid tenant field in response")
	}

	id, ok := tenantObj["id"].(string)
	if !ok {
		return nil, errors.New("Missing or invalid id field in tenant object")
	}

	name, ok := tenantObj["name"].(string)
	if !ok {
		return nil, errors.New("Missing or invalid name field in tenant object")
	}

	return map[string]string{
		"id":   id,
		"name": name,
	}, nil
}

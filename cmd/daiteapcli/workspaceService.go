package daiteapcli

import (
    "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func GetCurrentWorkspace() (map[string]string, error) {
    method := "GET"
    endpoint := "/account/tenant"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "false", "false")

    workspace := make(map[string]string)

    if err != nil {
        return workspace, err
    }

    responseObject := responseBody["tenant"].(map[string]interface{})

    workspace["id"] = responseObject["id"].(string)
    workspace["name"] = responseObject["name"].(string)
    
    return workspace, nil
}
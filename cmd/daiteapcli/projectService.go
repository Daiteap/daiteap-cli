package daiteapcli

import (
    "errors"

    "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func GetProjectID(name string) (string, error) {
    method := "GET"
    endpoint := "/projects"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

    if err != nil {
        return "", err
    } 
    
    for _, project := range responseBody["data"].([]interface{}) {
        projectObject := project.(map[string]interface{})
        if projectObject["name"] == name {
            return projectObject["id"].(string), nil
        }
    }
    
    return "", errors.New("Error getting project")
}
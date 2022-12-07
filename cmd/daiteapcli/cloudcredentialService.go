package daiteapcli

import (
    "time"

    "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func ValidateCredentials(provider string, credentials map[string]interface{}) (bool, error) {
    method := "POST"
    endpoint := "/cloud-credentials/validate"

    workspace, err := GetCurrentWorkspace()
    if err != nil {
        return false, err
    }

    requestBody := ""

    if provider == "google" {
        googleKey := credentials["google_key"].(string)
        requestBody = "{\"credentials\": {\"google\": {\"google_key\": \"" + googleKey + "\"}}, \"tenant_id\": \"" + workspace["id"] + "\"}"
    } else if provider == "aws" {
        keyID := credentials["aws_access_key_id"].(string)
        keySecret := credentials["aws_secret_access_key"].(string)
        requestBody = "{\"credentials\": {\"aws\": {\"aws_access_key_id\": \"" + keyID + "\", \"aws_secret_access_key\": \"" + keySecret + "\"}}, \"tenant_id\": \"" + workspace["id"] + "\"}"
    } else if provider == "azure" {
        tenantID := credentials["azure_tenant_id"].(string)
        subscriptionID := credentials["azure_subscription_id"].(string)
        clientID := credentials["azure_client_id"].(string)
        clientSecret := credentials["azure_client_secret"].(string)
        requestBody = "{\"credentials\": {\"azure\": {\"azure_tenant_id\": \"" + tenantID + "\", \"azure_subscription_id\": \"" + subscriptionID + "\", \"azure_client_id\": \"" + clientID + "\", \"azure_client_secret\": \"" + clientSecret + "\"}}, \"tenant_id\": \"" + workspace["id"] + "\"}"
    }
    
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody, "true", "false", "false")
    if err != nil {
        return false, err
    }

    method = "GET"
    endpoint = "/task-message/" + responseBody["taskId"].(string)
    
    for i := 0; i < 20; i++ {
        responseBody, err = daiteapcli.SendDaiteapRequest(method, endpoint, "", "false", "false", "false")
        if err != nil {
            return false, err
        }
        if responseBody["status"] == "SUCCESS" {
            return true, nil
        }
        time.Sleep(time.Second * 1)
    }

    return false, nil
}
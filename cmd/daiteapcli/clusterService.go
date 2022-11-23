package daiteapcli

import (
    "strings"

    "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func IsKubernetes(clusterID string) (bool, error) {
    method := "GET"
    endpoint := "/clusters"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", "false", "false")

    if err != nil {
        return false, err
    } else {
        for _, cluster := range responseBody["data"].([]interface{}) {
            clusterObject := cluster.(map[string]interface{})
            if clusterObject["id"] == clusterID {
                switch clusterObject["type"].(float64) {
                case
                    1,
                    3,
                    5,
                    7:
                    return true, nil
                }
            }
        }
    }
    
    return false, nil
}

func IsCompute(clusterID string) (bool, error) {
    method := "GET"
    endpoint := "/clusters"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", "false", "false")

    if err != nil {
        return false, err
    } else {
        for _, cluster := range responseBody["data"].([]interface{}) {
            clusterObject := cluster.(map[string]interface{})
            if clusterObject["id"] == clusterID {
                switch clusterObject["type"].(float64) {
                case
                    2,
                    6:
                    return true, nil
                }
            }
        }
    }
    
    return false, nil
}

func GetValidZones(provider string, credentialID string, region string) ([]string, error) {
    method := "GET"
	endpoint := "/cloud-credentials/" + credentialID + "/regions/" + region + "/zones"
	responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", "false", "false")

    var zones []string

    if err != nil {
        return zones, err
    } else {
        zones = make([]string, len(responseBody["zones"].([]interface{})))

        for index, zone := range responseBody["zones"].([]interface{}) {
            zones[index] = zone.(string)
        }
    }
    
    return zones, nil
}

func GetValidInstanceTypes(provider string, credentialID string, region string, zone string) (map[string]string, error) {
    method := "GET"
	endpoint := "/cloud-credentials/" + credentialID + "/regions/" + region + "/zones/" + zone + "/instances"
	responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", "false", "false")

    instances := make(map[string]string)

    if err != nil {
        return instances, err
    } else {
        for _, instance := range responseBody["instances"].([]interface{}) {
            instanceObject := instance.(map[string]interface{})
            if strings.Contains(instanceObject["description"].(string), "Small") {
                instances["S"] = instanceObject["name"].(string)
            } else if strings.Contains(instanceObject["description"].(string), "Medium") {
                instances["M"] = instanceObject["name"].(string)
            } else if strings.Contains(instanceObject["description"].(string), "XLarge") {
                instances["XL"] = instanceObject["name"].(string)
            } else if strings.Contains(instanceObject["description"].(string), "Large") {
                instances["L"] = instanceObject["name"].(string)
            }
        }
    }
    
    return instances, nil
}

func GetValidOperatingSystems(provider string, credentialID string, region string, environmentType string, username string) ([]string, error) {
    method := "GET"
    endpoint := "cloud-credentials/" + credentialID + "/regions/" + region + "/environment-type/" + environmentType + "/operating-systems"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", "false", "false")

    var operatingSystems []string

    if err != nil {
        return operatingSystems, err
    } else {
        operatingSystems = make([]string, len(responseBody["operatingSystems"].([]interface{})))

        for index, operatingSystem := range responseBody["operatingSystems"].([]interface{}) {
            operatingSystemObject := operatingSystem.(map[string]interface{})
            operatingSystems[index] = operatingSystemObject["value"].(string)
        }
    }
    
    return operatingSystems, nil
}

func GetSupportedKubernetesConfig() (map[string]interface{}, error) {
    method := "GET"
    endpoint := "/clusters/dlcmv2-supported-configurations"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "false", "false", "false")

    k8sConfigs := make(map[string]interface{})

    if err != nil {
        return k8sConfigs, err
    }
    
    return responseBody, nil
}
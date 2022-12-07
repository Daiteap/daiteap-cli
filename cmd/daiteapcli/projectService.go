package daiteapcli

import (
    "encoding/json"
    "errors"
    "strings"
    "fmt"

    "github.com/spf13/cobra"
    "github.com/rodaine/table"
    "github.com/fatih/color"
    "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
)

func GetProjectID(name string) (string, error) {
    method := "GET"
    endpoint := "/projects"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", "false", "false")

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

func ListProjectK8s(cmd *cobra.Command) () {
    verbose, _ := cmd.Flags().GetString("verbose")
    dryRun, _ := cmd.Flags().GetString("dry-run")
    outputFormat, _ := cmd.Flags().GetString("output")
    id, _ := cmd.Flags().GetString("id")
    name, _ := cmd.Flags().GetString("name")

    if len(id) <= 1 {
        id, _ = GetProjectID(name)
    }

    method := "GET"
	endpoint := "/clusters"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

    if err != nil {
        fmt.Println(err)
    } else if dryRun == "false" {
        clusterArray := make(map[string][]interface{})
        for _, cluster := range responseBody["data"].([]interface{}) {
            clusterObject := cluster.(map[string]interface{})
            switch clusterObject["type"].(float64) {
            case
                1,
                3,
                5,
                7:
                if clusterObject["project"] == id {
                    clusterArray["data"] = append(clusterArray["data"], clusterObject)
                }
            }
        }

        if outputFormat == "json" {
            output, _ := json.MarshalIndent(clusterArray, "", "    ")
            fmt.Println(string(output))
        } else if outputFormat == "wide" {
            tbl := table.New("ID", "Name", "Project", "Description", "Type", "Provider", "Created at", "Created by", "Status")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, cluster := range clusterArray["data"] {
                clusterObject := cluster.(map[string]interface{})

                providers := strings.ReplaceAll(clusterObject["providers"].(string), "[", "")
                providers = strings.ReplaceAll(providers, "]", "")
                providers = strings.ReplaceAll(providers, "\"", "")
                providersArray := strings.Split(providers, ",")
                providers = ""
                for _, provider := range providersArray {
                    if len(providers) == 0 {
                        providers += provider
                    } else {
                        providers += ", " + provider
                    }
                }

                environmentType := ""
                switch clusterObject["type"].(float64) {
                case 1:
                    environmentType = "DLCM"
                case 3:
                    environmentType = "DK3S"
                case 5:
                    environmentType = "CAPI"
                case 7:
                    environmentType = "DLCMv2"
                }

                installstep := clusterObject["installstep"].(float64)
                status := ""
                if installstep < 0 && installstep > -100 {
                    status = "Error creating"
                } else if installstep > 0 && installstep < 100 {
                    status = "Creating"
                } else if installstep == -100 {
                    status = "Error deleting"
                } else if installstep == 100 {
                    status = "Deleting"
                } else {
                    if clusterObject["status"].(float64) == 0 {
                        status = "Running"
                    } else if clusterObject["status"].(float64) == 1 {
                        status = "Starting"
                    } else if clusterObject["status"].(float64) == -1 {
                        status = "Error starting"
                    } else if clusterObject["status"].(float64) == 2 {
                        status = "Stopping"
                    } else if clusterObject["status"].(float64) == -2 {
                        status = "Error stopping"
                    } else if clusterObject["status"].(float64) == 3 {
                        status = "Restarting"
                    } else if clusterObject["status"].(float64) == -3 {
                        status = "Error restarting"
                    } else if clusterObject["status"].(float64) == 10 {
                        status = "Stopped"
                    }
                }

                tbl.AddRow(clusterObject["id"], clusterObject["name"], clusterObject["project_name"], clusterObject["description"], environmentType, providers,  clusterObject["created_at"], clusterObject["contact"], status)
            }

            tbl.Print()
        } else {
            tbl := table.New("Name", "Project", "Description", "Type", "Provider", "Created at", "Created by", "Status")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, cluster := range clusterArray["data"] {
                clusterObject := cluster.(map[string]interface{})

                providers := strings.ReplaceAll(clusterObject["providers"].(string), "[", "")
                providers = strings.ReplaceAll(providers, "]", "")
                providers = strings.ReplaceAll(providers, "\"", "")
                providersArray := strings.Split(providers, ",")
                providers = ""
                for _, provider := range providersArray {
                    if len(providers) == 0 {
                        providers += provider
                    } else {
                        providers += ", " + provider
                    }
                }

                environmentType := ""
                switch clusterObject["type"].(float64) {
                case 1:
                    environmentType = "DLCM"
                case 3:
                    environmentType = "DK3S"
                case 5:
                    environmentType = "CAPI"
                case 7:
                    environmentType = "DLCMv2"
                }

                installstep := clusterObject["installstep"].(float64)
                status := ""
                if installstep < 0 && installstep > -100 {
                    status = "Error creating"
                } else if installstep > 0 && installstep < 100 {
                    status = "Creating"
                } else if installstep == -100 {
                    status = "Error deleting"
                } else if installstep == 100 {
                    status = "Deleting"
                } else {
                    if clusterObject["status"].(float64) == 0 {
                        status = "Running"
                    } else if clusterObject["status"].(float64) == 1 {
                        status = "Starting"
                    } else if clusterObject["status"].(float64) == -1 {
                        status = "Error starting"
                    } else if clusterObject["status"].(float64) == 2 {
                        status = "Stopping"
                    } else if clusterObject["status"].(float64) == -2 {
                        status = "Error stopping"
                    } else if clusterObject["status"].(float64) == 3 {
                        status = "Restarting"
                    } else if clusterObject["status"].(float64) == -3 {
                        status = "Error restarting"
                    } else if clusterObject["status"].(float64) == 10 {
                        status = "Stopped"
                    }
                }

                tbl.AddRow(clusterObject["name"], clusterObject["project_name"], clusterObject["description"], environmentType, providers,  clusterObject["created_at"], clusterObject["contact"], status)
            }

            tbl.Print()
        }
    }
}

func ListProjectCompute(cmd *cobra.Command) () {
    verbose, _ := cmd.Flags().GetString("verbose")
    dryRun, _ := cmd.Flags().GetString("dry-run")
    outputFormat, _ := cmd.Flags().GetString("output")
    id, _ := cmd.Flags().GetString("id")
    name, _ := cmd.Flags().GetString("name")

    if len(id) <= 1 {
        id, _ = GetProjectID(name)
    }

    method := "GET"
	endpoint := "/clusters"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

    if err != nil {
        fmt.Println(err)
    } else if dryRun == "false" {
        clusterArray := make(map[string][]interface{})
        for _, cluster := range responseBody["data"].([]interface{}) {
            clusterObject := cluster.(map[string]interface{})
            switch clusterObject["type"].(float64) {
            case
                2,
                6:
                if clusterObject["project"] == id {
                    clusterArray["data"] = append(clusterArray["data"], clusterObject)
                }
            }
        }

        if outputFormat == "json" {
            output, _ := json.MarshalIndent(clusterArray, "", "    ")
            fmt.Println(string(output))
        } else if outputFormat == "wide" {
            tbl := table.New("ID", "Name", "Project", "Provider", "Created at", "Created by", "Status")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, cluster := range clusterArray["data"] {
                clusterObject := cluster.(map[string]interface{})

                providers := strings.ReplaceAll(clusterObject["providers"].(string), "[", "")
                providers = strings.ReplaceAll(providers, "]", "")
                providers = strings.ReplaceAll(providers, "\"", "")
                providersArray := strings.Split(providers, ",")
                providers = ""
                for _, provider := range providersArray {
                    if len(providers) == 0 {
                        providers += provider
                    } else {
                        providers += ", " + provider
                    }
                }

                installstep := clusterObject["installstep"].(float64)
                status := ""
                if installstep < 0 && installstep > -100 {
                    status = "Error creating"
                } else if installstep > 0 && installstep < 100 {
                    status = "Creating"
                } else if installstep == -100 {
                    status = "Error deleting"
                } else if installstep == 100 {
                    status = "Deleting"
                } else {
                    if clusterObject["status"].(float64) == 0 {
                        status = "Running"
                    } else if clusterObject["status"].(float64) == 1 {
                        status = "Starting"
                    } else if clusterObject["status"].(float64) == -1 {
                        status = "Error starting"
                    } else if clusterObject["status"].(float64) == 2 {
                        status = "Stopping"
                    } else if clusterObject["status"].(float64) == -2 {
                        status = "Error stopping"
                    } else if clusterObject["status"].(float64) == 3 {
                        status = "Restarting"
                    } else if clusterObject["status"].(float64) == -3 {
                        status = "Error restarting"
                    } else if clusterObject["status"].(float64) == 10 {
                        status = "Stopped"
                    }
                }

                tbl.AddRow(clusterObject["id"], clusterObject["name"], clusterObject["project_name"], providers,  clusterObject["created_at"], clusterObject["contact"], status)
            }

            tbl.Print()
        } else {
            tbl := table.New("Name", "Project", "Provider", "Created at", "Created by", "Status")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, cluster := range clusterArray["data"] {
                clusterObject := cluster.(map[string]interface{})

                providers := strings.ReplaceAll(clusterObject["providers"].(string), "[", "")
                providers = strings.ReplaceAll(providers, "]", "")
                providers = strings.ReplaceAll(providers, "\"", "")
                providersArray := strings.Split(providers, ",")
                providers = ""
                for _, provider := range providersArray {
                    if len(providers) == 0 {
                        providers += provider
                    } else {
                        providers += ", " + provider
                    }
                }

                installstep := clusterObject["installstep"].(float64)
                status := ""
                if installstep < 0 && installstep > -100 {
                    status = "Error creating"
                } else if installstep > 0 && installstep < 100 {
                    status = "Creating"
                } else if installstep == -100 {
                    status = "Error deleting"
                } else if installstep == 100 {
                    status = "Deleting"
                } else {
                    if clusterObject["status"].(float64) == 0 {
                        status = "Running"
                    } else if clusterObject["status"].(float64) == 1 {
                        status = "Starting"
                    } else if clusterObject["status"].(float64) == -1 {
                        status = "Error starting"
                    } else if clusterObject["status"].(float64) == 2 {
                        status = "Stopping"
                    } else if clusterObject["status"].(float64) == -2 {
                        status = "Error stopping"
                    } else if clusterObject["status"].(float64) == 3 {
                        status = "Restarting"
                    } else if clusterObject["status"].(float64) == -3 {
                        status = "Error restarting"
                    } else if clusterObject["status"].(float64) == 10 {
                        status = "Stopped"
                    }
                }

                tbl.AddRow(clusterObject["name"], clusterObject["project_name"], providers,  clusterObject["created_at"], clusterObject["contact"], status)
            }

            tbl.Print()
        }
    }
}


func ListProjectStorage(cmd *cobra.Command) () {
    verbose, _ := cmd.Flags().GetString("verbose")
    dryRun, _ := cmd.Flags().GetString("dry-run")
    outputFormat, _ := cmd.Flags().GetString("output")
    id, _ := cmd.Flags().GetString("id")
    name, _ := cmd.Flags().GetString("name")

    if len(id) <= 1 {
        id, _ = GetProjectID(name)
    }

    method := "GET"
    endpoint := "/buckets"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

    if err != nil {
        fmt.Println(err)
    } else if dryRun == "false" {
        storageArray := make(map[string][]interface{})
        for _, cluster := range responseBody["data"].([]interface{}) {
            storageObject := cluster.(map[string]interface{})
            switch storageObject["type"].(float64) {
            case
                2,
                6:
                if storageObject["project"] == id {
                    storageArray["data"] = append(storageArray["data"], storageObject)
                }
            }
        }

        if outputFormat == "json" {
            output, _ := json.MarshalIndent(storageArray, "", "    ")
            fmt.Println(string(output))
        } else if outputFormat == "wide" {
            tbl := table.New("ID", "Name", "Cloud", "Project", "Credential", "Created At")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, bucket := range storageArray["data"] {
                bucketObject := bucket.(map[string]interface{})

                tbl.AddRow(bucketObject["id"], bucketObject["name"], bucketObject["provider"], bucketObject["project"], bucketObject["credential"], bucketObject["created_at"])
            }

            tbl.Print()
        } else {
            tbl := table.New("Name", "Cloud", "Project", "Credential", "Created At")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, bucket := range storageArray["data"] {
                bucketObject := bucket.(map[string]interface{})

                tbl.AddRow(bucketObject["name"], bucketObject["provider"], bucketObject["project"], bucketObject["credential"], bucketObject["created_at"])
            }

            tbl.Print()
        }
    }
}


func ListProjectUsers(cmd *cobra.Command) () {
    verbose, _ := cmd.Flags().GetString("verbose")
    dryRun, _ := cmd.Flags().GetString("dry-run")
    outputFormat, _ := cmd.Flags().GetString("output")
    id, _ := cmd.Flags().GetString("id")
    name, _ := cmd.Flags().GetString("name")

    if len(id) <= 1 {
        id, _ = GetProjectID(name)
    }

    method := "GET"
    endpoint := "/projects/" + id + "/users"
    responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "", "true", verbose, dryRun)

    if err != nil {
        fmt.Println(err)
    } else if dryRun == "false" {
        if outputFormat == "json" {
            output, _ := json.MarshalIndent(responseBody, "", "    ")
            fmt.Println(string(output))
        } else if outputFormat == "wide" {
            tbl := table.New("ID", "User", "Role", "Projects", "Phone Number")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, user := range responseBody["users_list"].([]interface{}) {
                userObject := user.(map[string]interface{})
                projects := ""
                for _, project := range userObject["projects"].([]interface {}) {
                    if len(projects) == 0 {
                        projects += project.(string)
                    } else {
                        projects += ", " + project.(string)
                    }
                }

                tbl.AddRow(userObject["id"], userObject["username"], userObject["role"], projects, userObject["phone"])
            }

            tbl.Print()
        } else {
            tbl := table.New("User", "Role", "Projects", "Phone Number")
            headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
            columnFmt := color.New(color.FgYellow).SprintfFunc()
            tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

            for _, user := range responseBody["users_list"].([]interface{}) {
                userObject := user.(map[string]interface{})
                projects := ""
                for _, project := range userObject["projects"].([]interface {}) {
                    if len(projects) == 0 {
                        projects += project.(string)
                    } else {
                        projects += ", " + project.(string)
                    }
                }

                tbl.AddRow(userObject["username"], userObject["role"], projects, userObject["phone"])
            }

            tbl.Print()
        }
    }
}
package daiteapcli_test

import (
	"errors"
	"testing"

	daiteapcmd "github.com/Daiteap/daiteapcli/cmd/daiteapcli"
	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	daiteappkg "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func TestRunStorageDetailsCmd_Success(t *testing.T) {
	// Mock the DaiteapcliSendDaiteapRequest function
	sendDaiteapRequestCalledTimes := 0
	sendDaiteapRequestcalledMethod := ""
	sendDaiteapRequestcalledEndpoint := ""
	sendDaiteapRequestcalledRequestBody := ""
	sendDaiteapRequestcalledTenant := ""
	sendDaiteapRequestcalledVerbose := ""
	sendDaiteapRequestcalledDryRun := ""

	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		sendDaiteapRequestCalledTimes++
		sendDaiteapRequestcalledMethod = method
		sendDaiteapRequestcalledEndpoint = endpoint
		sendDaiteapRequestcalledRequestBody = requestBody
		sendDaiteapRequestcalledTenant = tenant
		sendDaiteapRequestcalledVerbose = verbose
		sendDaiteapRequestcalledDryRun = dryRun

		responseBody := map[string]interface{}{
			"activeTenants": []interface{}{
				map[string]interface{}{
					"id":        "1",
					"name":      "Workspace 1",
					"owner":     "John Doe",
					"email":     "john.doe@example.com",
					"phone":     "555-1234",
					"createdAt": "2022-01-01T00:00:00Z",
					"updatedAt": "2022-01-01T00:00:00Z",
					"selected":  true,
				},
				map[string]interface{}{
					"id":        "2",
					"name":      "Workspace 2",
					"owner":     "Jane Doe",
					"email":     "jane.doe@example.com",
					"phone":     "555-5678",
					"createdAt": "2022-01-02T00:00:00Z",
					"updatedAt": "2022-01-02T00:00:00Z",
					"selected":  false,
				},
			},
		}

		return responseBody, nil
	}

	// Mock the FmtPrintln function
	fmtPrintlnMessage := map[string]interface{}{
		"activeTenants": []interface{}{
			map[string]interface{}{
				"id":        "1",
				"name":      "Workspace 1",
				"owner":     "John Doe",
				"email":     "john.doe@example.com",
				"phone":     "555-1234",
				"createdAt": "2022-01-01T00:00:00Z",
				"updatedAt": "2022-01-01T00:00:00Z",
				"selected":  true,
			},
			map[string]interface{}{
				"id":        "2",
				"name":      "Workspace 2",
				"owner":     "Jane Doe",
				"email":     "jane.doe@example.com",
				"phone":     "555-5678",
				"createdAt": "2022-01-02T00:00:00Z",
				"updatedAt": "2022-01-02T00:00:00Z",
				"selected":  false,
			},
		},
	}

	expectedFmtPrintlnMessage, _ := daiteapcli.JsonMarshalIndent(fmtPrintlnMessage, "", "    ")

	fmtPrintlnCalledTimes := 0
	fmtPrintlnCalledMessage := ""
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledMessage = a[0].(string)
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Call the RunStorageDetailsCmd function
	daiteapcmd.RunStorageDetailsCmd(cmd, []string{})

	if fmtPrintlnCalledTimes != 1 {
		t.Errorf("expected FmtPrintln to be called once, but got %v", fmtPrintlnCalledTimes)
	}

	if fmtPrintlnCalledMessage != string(expectedFmtPrintlnMessage) {
		t.Errorf("expected FmtPrintln to be called with '%v', but got %v", fmtPrintlnCalledMessage, string(expectedFmtPrintlnMessage))
	}

	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	if sendDaiteapRequestcalledMethod != "GET" {
		t.Errorf("expected SendDaiteapRequest to be called with 'GET', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/buckets/" {
		t.Errorf("expected SendDaiteapRequest to be called with '/buckets/', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledRequestBody)
	}

	if sendDaiteapRequestcalledTenant != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledTenant)
	}

	if sendDaiteapRequestcalledVerbose != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledVerbose)
	}

	if sendDaiteapRequestcalledDryRun != "false" {
		t.Errorf("expected SendDaiteapRequest to be called with 'false', but got %v", sendDaiteapRequestcalledDryRun)
	}
}

func TestRunStorageDetailsCmd_Failure(t *testing.T) {
	// Mock the DaiteapcliSendDaiteapRequest function
	sendDaiteapRequestCalledTimes := 0
	sendDaiteapRequestcalledMethod := ""
	sendDaiteapRequestcalledEndpoint := ""
	sendDaiteapRequestcalledRequestBody := ""
	sendDaiteapRequestcalledTenant := ""
	sendDaiteapRequestcalledVerbose := ""
	sendDaiteapRequestcalledDryRun := ""

	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		sendDaiteapRequestCalledTimes++
		sendDaiteapRequestcalledMethod = method
		sendDaiteapRequestcalledEndpoint = endpoint
		sendDaiteapRequestcalledRequestBody = requestBody
		sendDaiteapRequestcalledTenant = tenant
		sendDaiteapRequestcalledVerbose = verbose
		sendDaiteapRequestcalledDryRun = dryRun

		return nil, errors.New("failed to send request")
	}

	// Mock the FmtPrintln function
	fmtPrintlnCalledTimes := 0
	fmtPrintlnCalledMessage := ""
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledMessage = a[0].(error).Error()
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Call the RunStorageDetailsCmd function
	daiteapcmd.RunStorageDetailsCmd(cmd, []string{})

	if fmtPrintlnCalledTimes != 1 {
		t.Errorf("expected FmtPrintln to be called once, but got %v", fmtPrintlnCalledTimes)
	}

	if fmtPrintlnCalledMessage != "failed to send request" {
		t.Errorf("expected FmtPrintln to be not called with 'failed to send request', but got %v", fmtPrintlnCalledMessage)
	}

	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	if sendDaiteapRequestcalledMethod != "GET" {
		t.Errorf("expected SendDaiteapRequest to be called with 'GET', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/buckets/" {
		t.Errorf("expected SendDaiteapRequest to be called with '/buckets/', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledRequestBody)
	}

	if sendDaiteapRequestcalledTenant != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledTenant)
	}

	if sendDaiteapRequestcalledVerbose != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledVerbose)
	}

	if sendDaiteapRequestcalledDryRun != "false" {
		t.Errorf("expected SendDaiteapRequest to be called with 'false', but got %v", sendDaiteapRequestcalledDryRun)
	}
}

func TestRunStorageDetailsCmd_Success_Dryrun(t *testing.T) {
	// Mock the DaiteapcliSendDaiteapRequest function
	sendDaiteapRequestCalledTimes := 0
	sendDaiteapRequestcalledMethod := ""
	sendDaiteapRequestcalledEndpoint := ""
	sendDaiteapRequestcalledRequestBody := ""
	sendDaiteapRequestcalledTenant := ""
	sendDaiteapRequestcalledVerbose := ""
	sendDaiteapRequestcalledDryRun := ""

	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		sendDaiteapRequestCalledTimes++
		sendDaiteapRequestcalledMethod = method
		sendDaiteapRequestcalledEndpoint = endpoint
		sendDaiteapRequestcalledRequestBody = requestBody
		sendDaiteapRequestcalledTenant = tenant
		sendDaiteapRequestcalledVerbose = verbose
		sendDaiteapRequestcalledDryRun = dryRun

		responseBody := map[string]interface{}{
			"activeTenants": []interface{}{
				map[string]interface{}{
					"id":        "1",
					"name":      "Workspace 1",
					"owner":     "John Doe",
					"email":     "john.doe@example.com",
					"phone":     "555-1234",
					"createdAt": "2022-01-01T00:00:00Z",
					"updatedAt": "2022-01-01T00:00:00Z",
					"selected":  true,
				},
				map[string]interface{}{
					"id":        "2",
					"name":      "Workspace 2",
					"owner":     "Jane Doe",
					"email":     "jane.doe@example.com",
					"phone":     "555-5678",
					"createdAt": "2022-01-02T00:00:00Z",
					"updatedAt": "2022-01-02T00:00:00Z",
					"selected":  false,
				},
			},
		}

		return responseBody, nil
	}

	// Mock the FmtPrintln function
	fmtPrintlnCalledTimes := 0
	fmtPrintlnCalledMessage := ""
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledMessage = a[0].(string)
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "true", "true")

	// Call the RunStorageDetailsCmd function
	daiteapcmd.RunStorageDetailsCmd(cmd, []string{})

	if fmtPrintlnCalledTimes != 0 {
		t.Errorf("expected FmtPrintln to not be called, but got %v", fmtPrintlnCalledTimes)
	}

	if fmtPrintlnCalledMessage != "" {
		t.Errorf("expected FmtPrintln to be called with '%v', but got %v", fmtPrintlnCalledMessage, "")
	}

	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	if sendDaiteapRequestcalledMethod != "GET" {
		t.Errorf("expected SendDaiteapRequest to be called with 'GET', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/buckets/" {
		t.Errorf("expected SendDaiteapRequest to be called with '/buckets/', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledRequestBody)
	}

	if sendDaiteapRequestcalledTenant != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledTenant)
	}

	if sendDaiteapRequestcalledVerbose != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledVerbose)
	}

	if sendDaiteapRequestcalledDryRun != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'false', but got %v", sendDaiteapRequestcalledDryRun)
	}
}

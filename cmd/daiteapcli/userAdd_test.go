package daiteapcli_test

import (
	"encoding/json"
	"errors"
	"testing"

	daiteapcmd "github.com/Daiteap/daiteapcli/cmd/daiteapcli"
	daiteappkg "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

func TestRunUserAddCmd_Success(t *testing.T) {
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
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, nil
	}

	fmtPrintlnCalledTimes := 0
	var fmtPrintlnCalledArgs []any = nil
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledArgs = a
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Call the RunUserAddCmd function
	daiteapcmd.RunUserAddCmd(cmd, []string{})

	// Check that the FmtPrintln function was called with the expected arguments
	fmtPrintlnExpectedCallArg, _ := json.MarshalIndent(map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, "", "    ")

	// Check that the FmtPrintln function was called
	if fmtPrintlnCalledTimes != 1 {
		t.Errorf("expected FmtPrintln to be called once, but got %v", fmtPrintlnCalledTimes)
	}

	// Check fmtPrintlnCalledArgs length
	if len(fmtPrintlnCalledArgs) != 1 {
		t.Errorf("expected FmtPrintln to be called with 1 argument, but got %v", len(fmtPrintlnCalledArgs))
	}

	if string(fmtPrintlnExpectedCallArg) != fmtPrintlnCalledArgs[0].(string) {
		t.Errorf("expected FmtPrintln to be called with %v, but got %v", fmtPrintlnExpectedCallArg, fmtPrintlnCalledArgs[0].(string))
	}

	// Check that the SendDaiteapRequest function was called
	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	// Check that the SendDaiteapRequest function was called with the expected arguments
	if sendDaiteapRequestcalledMethod != "POST" {
		t.Errorf("expected SendDaiteapRequest to be called with 'POST', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}" {
		t.Errorf("expected SendDaiteapRequest to be called with '{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}', but got %v", sendDaiteapRequestcalledRequestBody)
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

func TestRunUserAddCmd_RequestSenderError(t *testing.T) {
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
		return nil, errors.New("error sending request")
	}

	var fmtPrintlnCalledTimes = 0
	var fmtPrintlnCalledArgs []any = nil

	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledArgs = a
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Call the RunUserAddCmd function with an empty slice of strings
	daiteapcmd.RunUserAddCmd(cmd, []string{})

	// Check that the FmtPrintln function was called
	if fmtPrintlnCalledTimes == 0 {
		t.Errorf("expected FmtPrintln to be called once, but got %v", fmtPrintlnCalledTimes)
	}

	expectedCallArg := errors.New("error sending request")
	calledArg := fmtPrintlnCalledArgs[0]

	expectedCallArgStr := expectedCallArg.Error()
	calledArgStr := calledArg.(error).Error()

	if expectedCallArgStr != calledArgStr {
		t.Errorf("expected FmtPrintln to be called with 'error sending request', but got '%v'", fmtPrintlnCalledArgs[0])
	}

	// Check that the SendDaiteapRequest function was called
	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	// Check that the SendDaiteapRequest function was called with the expected arguments
	if sendDaiteapRequestcalledMethod != "POST" {
		t.Errorf("expected SendDaiteapRequest to be called with 'POST', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}" {
		t.Errorf("expected SendDaiteapRequest to be called with '{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}', but got %v", sendDaiteapRequestcalledRequestBody)
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

// Test the RunUserAddCmd function with the dry-run flag set
func TestRunUserAddCmd_DryRun(t *testing.T) {
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
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, nil
	}

	fmtPrintlnCalledTimes := 0
	var fmtPrintlnCalledArgs []any = nil
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledArgs = a
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "true", "true")

	// Call the RunUserAddCmd function
	daiteapcmd.RunUserAddCmd(cmd, []string{})

	// Check that the FmtPrintln function was called
	if fmtPrintlnCalledTimes != 0 {
		t.Errorf("expected FmtPrintln to be called 0 times, but got %v", fmtPrintlnCalledTimes)
	}

	// Check fmtPrintlnCalledArgs length
	if len(fmtPrintlnCalledArgs) != 0 {
		t.Errorf("expected FmtPrintln to be called with 0 arguments, but got %v", len(fmtPrintlnCalledArgs))
	}

	// Check that the SendDaiteapRequest function was called
	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	// Check that the SendDaiteapRequest function was called with the expected arguments
	if sendDaiteapRequestcalledMethod != "POST" {
		t.Errorf("expected SendDaiteapRequest to be called with 'POST', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}" {
		t.Errorf("expected SendDaiteapRequest to be called with '{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}', but got %v", sendDaiteapRequestcalledRequestBody)
	}

	if sendDaiteapRequestcalledTenant != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledTenant)
	}

	if sendDaiteapRequestcalledVerbose != "" {
		t.Errorf("expected SendDaiteapRequest to be called with '', but got %v", sendDaiteapRequestcalledVerbose)
	}

	if sendDaiteapRequestcalledDryRun != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledDryRun)
	}
}

// Test the RunUserAddCmd function with the verbose flag set
func TestRunUserAddCmd_Verbose(t *testing.T) {
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
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, nil
	}

	fmtPrintlnCalledTimes := 0
	var fmtPrintlnCalledArgs []any = nil
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledArgs = a
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Set Verbose to true
	cmd.Flags().String("verbose", "true", "true")

	// Call the RunUserAddCmd function
	daiteapcmd.RunUserAddCmd(cmd, []string{})

	// Check that the FmtPrintln function was called with the expected arguments
	fmtPrintlnExpectedCallArg, _ := json.MarshalIndent(map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, "", "    ")

	// Check that the FmtPrintln function was called
	if fmtPrintlnCalledTimes != 1 {
		t.Errorf("expected FmtPrintln to be called once, but got %v", fmtPrintlnCalledTimes)
	}

	// Check fmtPrintlnCalledArgs length
	if len(fmtPrintlnCalledArgs) != 1 {
		t.Errorf("expected FmtPrintln to be called with 1 argument, but got %v", len(fmtPrintlnCalledArgs))
	}

	if string(fmtPrintlnExpectedCallArg) != fmtPrintlnCalledArgs[0].(string) {
		t.Errorf("expected FmtPrintln to be called with %v, but got %v", fmtPrintlnExpectedCallArg, fmtPrintlnCalledArgs[0].(string))
	}

	// Check that the SendDaiteapRequest function was called
	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	// Check that the SendDaiteapRequest function was called with the expected arguments
	if sendDaiteapRequestcalledMethod != "POST" {
		t.Errorf("expected SendDaiteapRequest to be called with 'POST', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}" {
		t.Errorf("expected SendDaiteapRequest to be called with '{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}', but got %v", sendDaiteapRequestcalledRequestBody)
	}

	if sendDaiteapRequestcalledTenant != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledTenant)
	}

	if sendDaiteapRequestcalledVerbose != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledVerbose)
	}

	if sendDaiteapRequestcalledDryRun != "false" {
		t.Errorf("expected SendDaiteapRequest to be called with 'false', but got %v", sendDaiteapRequestcalledDryRun)
	}
}

// Test the RunUserAddCmd function with both the dry-run and verbose flags set
func TestRunUserAddCmd_DryRunVerbose(t *testing.T) {
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
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, nil
	}

	fmtPrintlnCalledTimes := 0
	var fmtPrintlnCalledArgs []any = nil
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledArgs = a
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to true
	cmd.Flags().String("dry-run", "true", "true")

	// Set Verbose to true
	cmd.Flags().String("verbose", "true", "true")

	// Call the RunUserAddCmd function
	daiteapcmd.RunUserAddCmd(cmd, []string{})

	// Check that the FmtPrintln function was called
	if fmtPrintlnCalledTimes != 0 {
		t.Errorf("expected FmtPrintln to not be called, but got %v", fmtPrintlnCalledTimes)
	}

	// Check fmtPrintlnCalledArgs length
	if len(fmtPrintlnCalledArgs) != 0 {
		t.Errorf("expected FmtPrintln to be called with 0 arguments, but got %v", len(fmtPrintlnCalledArgs))
	}

	// Check that the SendDaiteapRequest function was called
	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	// Check that the SendDaiteapRequest function was called with the expected arguments
	if sendDaiteapRequestcalledMethod != "POST" {
		t.Errorf("expected SendDaiteapRequest to be called with 'POST', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}" {
		t.Errorf("expected SendDaiteapRequest to be called with '{\"username\": \"\", \"firstname\": \"\", \"lastname\": \"\", \"email\": \"\", \"company\": \"\", \"phone\": \"\", \"sshpubkey\": \"\", \"userRole\": \"\"}', but got %v", sendDaiteapRequestcalledRequestBody)
	}

	if sendDaiteapRequestcalledTenant != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledTenant)
	}

	if sendDaiteapRequestcalledVerbose != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledVerbose)
	}

	if sendDaiteapRequestcalledDryRun != "true" {
		t.Errorf("expected SendDaiteapRequest to be called with 'true', but got %v", sendDaiteapRequestcalledDryRun)
	}
}

func TestRunUserAddCmd_Success_changed_fields(t *testing.T) {
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
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, nil
	}

	fmtPrintlnCalledTimes := 0
	var fmtPrintlnCalledArgs []any = nil
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		fmtPrintlnCalledArgs = a
		return 0, nil
	}

	// Create a new command object
	cmd := &cobra.Command{}

	cmd.Flags().String("dry-run", "false", "false")
	cmd.Flags().String("username", "testuser", "testuser")
	cmd.Flags().String("firstname", "testfirstname", "testfirstname")
	cmd.Flags().String("lastname", "testlastname", "testlastname")
	cmd.Flags().String("email", "testemail", "testemail")
	cmd.Flags().String("company", "testcompany", "testcompany")
	cmd.Flags().String("phone", "testphone", "testphone")
	cmd.Flags().String("sshpubkey", "testsshpubkey", "testsshpubkey")
	cmd.Flags().String("user-role", "testuserrole", "testuserrole")

	// Call the RunUserAddCmd function
	daiteapcmd.RunUserAddCmd(cmd, []string{})

	// Check that the FmtPrintln function was called with the expected arguments
	fmtPrintlnExpectedCallArg, _ := json.MarshalIndent(map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "test"}}, "", "    ")

	// Check that the FmtPrintln function was called
	if fmtPrintlnCalledTimes != 1 {
		t.Errorf("expected FmtPrintln to be called once, but got %v", fmtPrintlnCalledTimes)
	}

	// Check fmtPrintlnCalledArgs length
	if len(fmtPrintlnCalledArgs) != 1 {
		t.Errorf("expected FmtPrintln to be called with 1 argument, but got %v", len(fmtPrintlnCalledArgs))
	}

	if string(fmtPrintlnExpectedCallArg) != fmtPrintlnCalledArgs[0].(string) {
		t.Errorf("expected FmtPrintln to be called with %v, but got %v", fmtPrintlnExpectedCallArg, fmtPrintlnCalledArgs[0].(string))
	}

	// Check that the SendDaiteapRequest function was called
	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	// Check that the SendDaiteapRequest function was called with the expected arguments
	if sendDaiteapRequestcalledMethod != "POST" {
		t.Errorf("expected SendDaiteapRequest to be called with 'POST', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
	}

	if sendDaiteapRequestcalledRequestBody != "{\"username\": \"testuser\", \"firstname\": \"testfirstname\", \"lastname\": \"testlastname\", \"email\": \"testemail\", \"company\": \"testcompany\", \"phone\": \"testphone\", \"sshpubkey\": \"testsshpubkey\", \"userRole\": \"testuserrole\"}" {
		t.Errorf("expected SendDaiteapRequest to be called with '{\"username\": \"testuser\", \"firstname\": \"testfirstname\", \"lastname\": \"testlastname\", \"email\": \"testemail\", \"company\": \"testcompany\", \"phone\": \"testphone\", \"sshpubkey\": \"testsshpubkey\", \"userRole\": \"testuserrole\"}', but got %v", sendDaiteapRequestcalledRequestBody)
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

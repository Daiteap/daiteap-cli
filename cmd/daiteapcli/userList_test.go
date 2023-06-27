package daiteapcli_test

import (
	"fmt"
	"testing"

	daiteapcmd "github.com/Daiteap/daiteapcli/cmd/daiteapcli"
	daiteappkg "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func TestRunUserListCmd_Success(t *testing.T) {
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
			"users_list": []interface{}{
				map[string]interface{}{
					"id":       1,
					"username": "johndoe",
					"role":     "admin",
					"projects": []interface{}{"project1", "project2"},
					"phone":    "555-1234",
				},
				map[string]interface{}{
					"id":       2,
					"username": "janedoe",
					"role":     "user",
					"projects": []interface{}{"project2", "project3"},
					"phone":    "555-5678",
				},
			},
		}

		return responseBody, nil
	}

	// Mock the FmtPrintln function
	fmtPrintlnCalledTimes := 0
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		return 0, nil
	}

	// Mock the TableNew function
	tableNewCalledTimes := 0
	var tableNewColumnHeaders = []interface{}{}
	var expectedTableNewColumnHeaders = []string{"User", "Role", "Projects", "Phone Number"}

	daiteappkg.TableNew = func(columnHeaders ...interface{}) table.Table {
		tableNewCalledTimes++
		tableNewColumnHeaders = columnHeaders

		return table.New(tableNewColumnHeaders)
	}

	// Mock the TablePrint function
	tablePrintCalledTimes := 0
	daiteappkg.TablePrint = func(t table.Table) {
		tablePrintCalledTimes++
	}

	// Mock the TableAddRow function
	tableAddRowCalledTimes := 0
	var tableAddRowCalledRows = []interface{}{}
	var expectedTableAddRowCalledRow = []interface{}{
		[]string{"Workspace 1", "John Doe", "john.doe@example.com", "555-1234", "2022-01-01T00:00:00Z", "2022-01-01T00:00:00Z", "true"},
		[]string{"Workspace 2", "Jane Doe", "jane.doe@example.com", "555-5678", "2022-01-02T00:00:00Z", "2022-01-02T00:00:00Z", "false"},
	}

	daiteappkg.TableAddRow = func(t table.Table, row ...interface{}) {
		tableAddRowCalledTimes++
		tableAddRowCalledRows = append(tableAddRowCalledRows, row)

	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")

	// Call the RunUserListCmd function
	daiteapcmd.RunUserListCmd(cmd, []string{})

	fmt.Println(tableNewColumnHeaders)
	for i, v := range tableNewColumnHeaders {
		vStr := fmt.Sprintf("%v", v)
		expectedTableNewColumnHeadersStr := fmt.Sprintf("%v", expectedTableNewColumnHeaders[i])

		if vStr != expectedTableNewColumnHeadersStr {
			t.Errorf("expected TableNew to be called with %v, but got %v", expectedTableNewColumnHeaders, tableNewColumnHeaders)
		}
	}

	for i, v := range tableAddRowCalledRows {
		vStr := fmt.Sprintf("%v", v)
		expectedTableAddRowCalledRowStr := fmt.Sprintf("%v", expectedTableAddRowCalledRow[i])
		if vStr != expectedTableAddRowCalledRowStr {
			t.Errorf("expected TableAddRow to be called with %v, but got %v", expectedTableAddRowCalledRow, tableAddRowCalledRows)
		}
	}

	if tableNewCalledTimes != 1 {
		t.Errorf("expected TableNew to be called once, but got %v", tableNewCalledTimes)
	}

	if tablePrintCalledTimes != 1 {
		t.Errorf("expected TablePrint to be called once, but got %v", tablePrintCalledTimes)
	}

	if fmtPrintlnCalledTimes != 0 {
		t.Errorf("expected FmtPrintln to be not called, but got %v", fmtPrintlnCalledTimes)
	}

	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	if sendDaiteapRequestcalledMethod != "GET" {
		t.Errorf("expected SendDaiteapRequest to be called with 'GET', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
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

func TestRunUserListCmd_OutputFormat_Wide(t *testing.T) {
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
			"users_list": []interface{}{
				map[string]interface{}{
					"id":       1,
					"username": "johndoe",
					"role":     "admin",
					"projects": []interface{}{"project1", "project2"},
					"phone":    "555-1234",
				},
				map[string]interface{}{
					"id":       2,
					"username": "janedoe",
					"role":     "user",
					"projects": []interface{}{"project2", "project3"},
					"phone":    "555-5678",
				},
			},
		}

		return responseBody, nil
	}

	// Mock the FmtPrintln function
	fmtPrintlnCalledTimes := 0
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		return 0, nil
	}

	// Mock the TableNew function
	tableNewCalledTimes := 0
	var tableNewColumnHeaders = []interface{}{}
	var expectedTableNewColumnHeaders = []string{"ID", "User", "Role", "Projects", "Phone Number"}

	daiteappkg.TableNew = func(columnHeaders ...interface{}) table.Table {
		tableNewCalledTimes++
		tableNewColumnHeaders = columnHeaders

		return table.New(tableNewColumnHeaders)
	}

	// Mock the TableAddRow function
	tableAddRowCalledTimes := 0
	var tableAddRowCalledRows = []interface{}{}
	var expectedTableAddRowCalledRow = []interface{}{
		[]string{"1", "Workspace 1", "John Doe", "john.doe@example.com", "555-1234", "2022-01-01T00:00:00Z", "2022-01-01T00:00:00Z", "true"},
		[]string{"2", "Workspace 2", "Jane Doe", "jane.doe@example.com", "555-5678", "2022-01-02T00:00:00Z", "2022-01-02T00:00:00Z", "false"},
	}

	daiteappkg.TableAddRow = func(t table.Table, row ...interface{}) {
		tableAddRowCalledTimes++
		tableAddRowCalledRows = append(tableAddRowCalledRows, row)

	}

	// Mock the TablePrint function
	tablePrintCalledTimes := 0
	daiteappkg.TablePrint = func(t table.Table) {
		tablePrintCalledTimes++
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")
	cmd.Flags().String("output", "wide", "wide")

	// Call the RunUserListCmd function
	daiteapcmd.RunUserListCmd(cmd, []string{})

	for i, v := range tableNewColumnHeaders {
		vStr := fmt.Sprintf("%v", v)
		expectedTableNewColumnHeadersStr := fmt.Sprintf("%v", expectedTableNewColumnHeaders[i])
		if vStr != expectedTableNewColumnHeadersStr {
			t.Errorf("expected TableNew to be called with %v, but got %v", expectedTableNewColumnHeaders, tableNewColumnHeaders)
		}
	}

	if tableNewCalledTimes != 1 {
		t.Errorf("expected TableNew to be called once, but got %v", tableNewCalledTimes)
	}

	for i, v := range tableAddRowCalledRows {
		vStr := fmt.Sprintf("%v", v)
		expectedTableAddRowCalledRowStr := fmt.Sprintf("%v", expectedTableAddRowCalledRow[i])
		if vStr != expectedTableAddRowCalledRowStr {
			t.Errorf("expected TableAddRow to be called with %v, but got %v", expectedTableAddRowCalledRow, tableAddRowCalledRows)
		}
	}

	if tablePrintCalledTimes != 1 {
		t.Errorf("expected TablePrint to be called once, but got %v", tablePrintCalledTimes)
	}

	if fmtPrintlnCalledTimes != 0 {
		t.Errorf("expected FmtPrintln to be not called, but got %v", fmtPrintlnCalledTimes)
	}

	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	if sendDaiteapRequestcalledMethod != "GET" {
		t.Errorf("expected SendDaiteapRequest to be called with 'GET', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
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

func TestRunUserListCmd_OutputFormat_Json(t *testing.T) {
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
			"users_list": []interface{}{
				map[string]interface{}{
					"id":       1,
					"username": "johndoe",
					"role":     "admin",
					"projects": []interface{}{"project1", "project2"},
					"phone":    "555-1234",
				},
				map[string]interface{}{
					"id":       2,
					"username": "janedoe",
					"role":     "user",
					"projects": []interface{}{"project2", "project3"},
					"phone":    "555-5678",
				},
			},
		}

		return responseBody, nil
	}

	// Mock the FmtPrintln function
	fmtPrintlnCalledTimes := 0
	daiteappkg.FmtPrintln = func(a ...any) (n int, err error) {
		fmtPrintlnCalledTimes++
		return 0, nil
	}

	// Mock the TableNew function
	tableNewCalledTimes := 0
	var tableNewColumnHeaders = []interface{}{}
	var expectedTableNewColumnHeaders = []string{"ID", "Name", "Owner", "Email", "Phone", "Created at", "Updated at", "Active"}

	daiteappkg.TableNew = func(columnHeaders ...interface{}) table.Table {
		tableNewCalledTimes++
		tableNewColumnHeaders = columnHeaders

		return table.New(tableNewColumnHeaders)
	}

	// Mock the TableAddRow function
	tableAddRowCalledTimes := 0
	var tableAddRowCalledRows = []interface{}{}
	var expectedTableAddRowCalledRow = []interface{}{
		[]string{"1", "Workspace 1", "John Doe", "john.doe@example.com", "555-1234", "2022-01-01T00:00:00Z", "2022-01-01T00:00:00Z", "true"},
		[]string{"2", "Workspace 2", "Jane Doe", "jane.doe@example.com", "555-5678", "2022-01-02T00:00:00Z", "2022-01-02T00:00:00Z", "false"},
	}

	daiteappkg.TableAddRow = func(t table.Table, row ...interface{}) {
		tableAddRowCalledTimes++
		tableAddRowCalledRows = append(tableAddRowCalledRows, row)

	}

	// Mock the TablePrint function
	tablePrintCalledTimes := 0
	daiteappkg.TablePrint = func(t table.Table) {
		tablePrintCalledTimes++
	}

	// Create a new command object
	cmd := &cobra.Command{}

	// Set DryRun to false
	cmd.Flags().String("dry-run", "false", "false")
	cmd.Flags().String("output", "json", "json")

	// Call the RunUserListCmd function
	daiteapcmd.RunUserListCmd(cmd, []string{})

	for i, v := range tableNewColumnHeaders {
		vStr := fmt.Sprintf("%v", v)
		expectedTableNewColumnHeadersStr := fmt.Sprintf("%v", expectedTableNewColumnHeaders[i])
		if vStr != expectedTableNewColumnHeadersStr {
			t.Errorf("expected TableNew to be called with %v, but got %v", expectedTableNewColumnHeaders, tableNewColumnHeaders)
		}
	}

	if tableNewCalledTimes != 0 {
		t.Errorf("expected TableNew to not be called, but got %v", tableNewCalledTimes)
	}

	for i, v := range tableAddRowCalledRows {
		vStr := fmt.Sprintf("%v", v)
		expectedTableAddRowCalledRowStr := fmt.Sprintf("%v", expectedTableAddRowCalledRow[i])
		if vStr != expectedTableAddRowCalledRowStr {
			t.Errorf("expected TableAddRow to be called with %v, but got %v", expectedTableAddRowCalledRow, tableAddRowCalledRows)
		}
	}

	if tablePrintCalledTimes != 0 {
		t.Errorf("expected TablePrint to not be called, but got %v", tablePrintCalledTimes)
	}

	if fmtPrintlnCalledTimes != 1 {
		t.Errorf("expected FmtPrintln to be called 1 time, but got %v", fmtPrintlnCalledTimes)
	}

	if sendDaiteapRequestCalledTimes != 1 {
		t.Errorf("expected SendDaiteapRequest to be called once, but got %v", sendDaiteapRequestCalledTimes)
	}

	if sendDaiteapRequestcalledMethod != "GET" {
		t.Errorf("expected SendDaiteapRequest to be called with 'GET', but got %v", sendDaiteapRequestcalledMethod)
	}

	if sendDaiteapRequestcalledEndpoint != "/users" {
		t.Errorf("expected SendDaiteapRequest to be called with '/users', but got %v", sendDaiteapRequestcalledEndpoint)
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

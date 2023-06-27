package daiteapcli

import (
	"errors"
	"testing"

	daiteappkg "github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentWorkspace_SuccessfulResponse(t *testing.T) {
	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123", "name": "Test Workspace"}}, nil
	}

	expectedWorkspace := map[string]string{
		"id":   "123",
		"name": "Test Workspace",
	}
	actualWorkspace, err := GetCurrentWorkspace()
	assert.NoError(t, err)
	assert.Equal(t, expectedWorkspace, actualWorkspace)
}

func TestGetCurrentWorkspace_ErrorResponse(t *testing.T) {
	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		return nil, errors.New("Failed to get workspace")
	}

	expectedWorkspace := map[string]string(nil)
	actualWorkspace, err := GetCurrentWorkspace()
	assert.Error(t, err)
	assert.Equal(t, expectedWorkspace, actualWorkspace)
}

func TestGetCurrentWorkspace_MissingTenantField(t *testing.T) {
	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		return map[string]interface{}{"foo": "bar"}, nil
	}

	expectedWorkspace := map[string]string(nil)
	actualWorkspace, err := GetCurrentWorkspace()
	assert.Error(t, err)
	assert.Equal(t, expectedWorkspace, actualWorkspace)
	assert.Contains(t, err.Error(), "Missing or invalid tenant field in response")
}

func TestGetCurrentWorkspace_MissingIdField(t *testing.T) {
	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		return map[string]interface{}{"tenant": map[string]interface{}{"name": "Test Workspace"}}, nil
	}

	expectedWorkspace := map[string]string(nil)
	actualWorkspace, err := GetCurrentWorkspace()
	assert.Error(t, err)
	assert.Equal(t, expectedWorkspace, actualWorkspace)
	assert.Contains(t, err.Error(), "Missing or invalid id field in tenant object")
}

func TestGetCurrentWorkspace_MissingNameField(t *testing.T) {
	daiteappkg.DaiteapcliSendDaiteapRequest = func(method string, endpoint string, requestBody string, tenant string, verbose string, dryRun string) (map[string]interface{}, error) {
		return map[string]interface{}{"tenant": map[string]interface{}{"id": "123"}}, nil
	}

	expectedWorkspace := map[string]string(nil)
	actualWorkspace, err := GetCurrentWorkspace()
	assert.Error(t, err)
	assert.Equal(t, expectedWorkspace, actualWorkspace)
	assert.Contains(t, err.Error(), "Missing or invalid name field in tenant object")
}

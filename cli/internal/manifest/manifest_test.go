package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStackNameFromManifest(t *testing.T) {
	// Create cdk.out directory
	err := os.MkdirAll("cdk.out", 0755)
	if err != nil {
		t.Fatalf("Failed to create cdk.out directory: %v", err)
	}
	defer func() {
		if err := os.RemoveAll("cdk.out"); err != nil {
			t.Errorf("Failed to clean up cdk.out directory: %v", err)
		}
	}()

	// Create manifest.json with test data
	manifestData := map[string]interface{}{
		"artifacts": map[string]interface{}{
			"test-stack": map[string]interface{}{
				"type": "aws:cloudformation:stack",
			},
		},
	}

	manifestBytes, err := json.Marshal(manifestData)
	if err != nil {
		t.Fatalf("Failed to marshal manifest data: %v", err)
	}

	manifestPath := filepath.Join("cdk.out", "manifest.json")
	if err := os.WriteFile(manifestPath, manifestBytes, 0644); err != nil {
		t.Fatalf("Failed to write manifest file: %v", err)
	}

	// Test successful case
	stackName, err := GetStackNameFromManifest()
	assert.NoError(t, err)
	assert.Equal(t, "test-stack", stackName)

	// Test case when stack is not found
	manifestData = map[string]interface{}{
		"artifacts": map[string]interface{}{
			"other-artifact": map[string]interface{}{
				"type": "other-type",
			},
		},
	}
	manifestBytes, err = json.Marshal(manifestData)
	if err != nil {
		t.Fatalf("Failed to marshal manifest data: %v", err)
	}

	if err := os.WriteFile(manifestPath, manifestBytes, 0644); err != nil {
		t.Fatalf("Failed to write manifest file: %v", err)
	}

	_, err = GetStackNameFromManifest()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no CloudFormation stack found in manifest.json")

	// Test case when manifest file does not exist
	if err := os.Remove(manifestPath); err != nil {
		t.Fatalf("Failed to remove manifest file: %v", err)
	}

	_, err = GetStackNameFromManifest()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read manifest.json")
}

func TestGetStackNameFromManifest_NoStack(t *testing.T) {
	// Create cdk.out directory
	err := os.MkdirAll("cdk.out", 0755)
	if err != nil {
		t.Fatalf("Failed to create cdk.out directory: %v", err)
	}
	defer func() {
		if err := os.RemoveAll("cdk.out"); err != nil {
			t.Errorf("Failed to clean up cdk.out directory: %v", err)
		}
	}()

	// Create manifest.json without stack
	manifestData := map[string]interface{}{
		"artifacts": map[string]interface{}{
			"other-artifact": map[string]interface{}{
				"type": "other-type",
			},
		},
	}

	manifestBytes, err := json.Marshal(manifestData)
	if err != nil {
		t.Fatalf("Failed to marshal manifest data: %v", err)
	}

	manifestPath := filepath.Join("cdk.out", "manifest.json")
	if err := os.WriteFile(manifestPath, manifestBytes, 0644); err != nil {
		t.Fatalf("Failed to write manifest file: %v", err)
	}

	// Run test
	_, err = GetStackNameFromManifest()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no CloudFormation stack found in manifest.json")
}

func TestGetStackNameFromManifest_NoFile(t *testing.T) {
	// Create and immediately remove cdk.out directory to ensure it doesn't exist
	err := os.MkdirAll("cdk.out", 0755)
	if err != nil {
		t.Fatalf("Failed to create cdk.out directory: %v", err)
	}
	if err := os.RemoveAll("cdk.out"); err != nil {
		t.Fatalf("Failed to remove cdk.out directory: %v", err)
	}

	// Run test
	_, err = GetStackNameFromManifest()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read manifest.json")
} 
package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetStackNameFromManifest(t *testing.T) {
	// Create test manifest.json
	testManifest := `{
		"artifacts": {
			"test-stack": {
				"type": "aws:cloudformation:stack"
			},
			"other-artifact": {
				"type": "other-type"
			}
		}
	}`

	// Create test directory
	testDir := "cdk.out"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create manifest.json
	manifestPath := filepath.Join(testDir, "manifest.json")
	err = os.WriteFile(manifestPath, []byte(testManifest), 0644)
	if err != nil {
		t.Fatalf("Failed to create test manifest.json: %v", err)
	}

	// Run test
	stackName, err := GetStackNameFromManifest()
	if err != nil {
		t.Fatalf("GetStackNameFromManifest failed: %v", err)
	}

	if stackName != "test-stack" {
		t.Errorf("Expected stack name 'test-stack', got '%s'", stackName)
	}
}

func TestGetStackNameFromManifest_NoStack(t *testing.T) {
	// Create manifest.json without stack
	testManifest := `{
		"artifacts": {
			"other-artifact": {
				"type": "other-type"
			}
		}
	}`

	// Create test directory
	testDir := "cdk.out"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create manifest.json
	manifestPath := filepath.Join(testDir, "manifest.json")
	err = os.WriteFile(manifestPath, []byte(testManifest), 0644)
	if err != nil {
		t.Fatalf("Failed to create test manifest.json: %v", err)
	}

	// Run test
	_, err = GetStackNameFromManifest()
	if err == nil {
		t.Error("Expected error when no stack is found, got nil")
	}
}

func TestGetStackNameFromManifest_NoFile(t *testing.T) {
	// Create test directory
	testDir := "cdk.out"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Run test
	_, err = GetStackNameFromManifest()
	if err == nil {
		t.Error("Expected error when manifest.json is not found, got nil")
	}
} 
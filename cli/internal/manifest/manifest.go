package manifest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Manifest struct {
	Artifacts map[string]Artifact `json:"artifacts"`
}

type Artifact struct {
	Type string `json:"type"`
}

// GetStackNameFromManifest retrieves the stack name from manifest.json in the cdk.out directory
func GetStackNameFromManifest() (string, error) {
	manifestPath := filepath.Join("cdk.out", "manifest.json")
	
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", fmt.Errorf("failed to read manifest.json: %w", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return "", fmt.Errorf("failed to parse manifest.json: %w", err)
	}

	// Find CloudFormation stack
	for name, artifact := range manifest.Artifacts {
		if artifact.Type == "aws:cloudformation:stack" {
			return name, nil
		}
	}

	return "", fmt.Errorf("no CloudFormation stack found in manifest.json")
} 
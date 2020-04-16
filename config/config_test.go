package config

import (
	"os"
	"testing"
)

func TestGetPath(t *testing.T) {
	expectedOutput := "config/config-prod.json"

	os.Setenv("ENVIRONMENT", "PROD")
	os.Setenv("PROD_CONFIG_PATH", expectedOutput)

	configPath := GetPath()

	if configPath != expectedOutput {
		t.Fatal("unexpected output")
	}
}

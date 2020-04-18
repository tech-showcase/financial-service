package config

import (
	"github.com/tech-showcase/financial-service/presenter"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedOutput := presenter.Config{
		DigitalCurrency: presenter.DigitalCurrency{
			ServerAddress: "http://dummy.address/",
			ApiKey:        "dummy-key",
		},
		Auth: presenter.Auth{
			ServerAddress: "http://dummy.address/",
		},
	}

	configPath := "config-example.json"
	os.Setenv("DEV_CONFIG_PATH", configPath)

	config, err := Parse()

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func TestGetPath(t *testing.T) {
	expectedOutput := "config/config-prod.json"

	os.Setenv("ENVIRONMENT", "PROD")
	os.Setenv("PROD_CONFIG_PATH", expectedOutput)

	configPath := GetPath()

	if configPath != expectedOutput {
		t.Fatal("unexpected output")
	}
}

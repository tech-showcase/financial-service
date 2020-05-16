package config

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	setDummyEnvVar()
	expectedOutput := getDummyConfig()

	config := Read()

	if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func setDummyEnvVar() {
	dummyConfig := getDummyConfig()

	os.Setenv("DC_SERVER_ADDRESS", dummyConfig.DigitalCurrency.ServerAddress)
	os.Setenv("DC_API_KEY", dummyConfig.DigitalCurrency.ApiKey)
	os.Setenv("AUTH_SERVER_ADDRESS", dummyConfig.Auth.ServerAddress)
}

func getDummyConfig() Config {
	dummyConfig := Config{
		DigitalCurrency: DigitalCurrency{
			ServerAddress: "http://dummyAddress",
			ApiKey:        "dummyApiKey",
		},
		Auth: Auth{
			ServerAddress: "http://dummyAddress",
		},
	}

	return dummyConfig
}

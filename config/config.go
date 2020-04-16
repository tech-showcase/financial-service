package config

import (
	"os"
)

func GetPath() string {
	environment := "DEV"
	if environmentFromEnvVar := os.Getenv("ENVIRONMENT"); environmentFromEnvVar != "" {
		environment = environmentFromEnvVar
	}

	configPath := "config/config.json"
	if configPathFromEnvVar := os.Getenv(environment + "_CONFIG_PATH"); configPathFromEnvVar != "" {
		configPath = configPathFromEnvVar
	}

	return configPath
}

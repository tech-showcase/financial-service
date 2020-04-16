package config

import (
	"encoding/json"
	"github.com/tech-showcase/financial-service/presenter"
	"io/ioutil"
	"os"
)

func Parse() (config presenter.Config, err error) {
	configPath := GetPath()

	configFileContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(configFileContent, &config)
	if err != nil {
		return
	}

	return
}

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

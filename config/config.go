package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	Config struct {
		DigitalCurrency DigitalCurrency `json:"digital_currency"`
		Auth            Auth            `json:"auth"`
	}

	DigitalCurrency struct {
		ServerAddress string `json:"server_address"`
		ApiKey        string `json:"api_key"`
	}

	Auth struct {
		ServerAddress string `json:"server_address"`
	}
)

func Read() (config Config, err error) {
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

	configPath := "config/config-dev.json"
	if configPathFromEnvVar := os.Getenv(environment + "_CONFIG_PATH"); configPathFromEnvVar != "" {
		configPath = configPathFromEnvVar
	}

	return configPath
}

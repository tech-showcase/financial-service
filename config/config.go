package config

import (
	"errors"
	"fmt"
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

const (
	EnvVarNotFoundErr = "env var %s not found"
)

var Instance = Config{}

func Read() (config Config, err error) {
	config, err = readFromEnvVar()

	return
}

func readFromEnvVar() (config Config, err error) {
	config.DigitalCurrency.ServerAddress = readEnvVarWithDefaultValue("DC_SERVER_ADDRESS", "http://localhost")
	config.DigitalCurrency.ApiKey, err = readMandatoryEnvVar("DC_API_KEY")

	config.Auth.ServerAddress = readEnvVarWithDefaultValue("AUTH_SERVER_ADDRESS", "http://localhost")

	return
}

func readEnvVarWithDefaultValue(key, defaultValue string) string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue
	}
	return defaultValue
}

func readMandatoryEnvVar(key string) (string, error) {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue, nil
	}
	return "", errors.New(fmt.Sprintf(EnvVarNotFoundErr, key))
}

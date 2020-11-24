package config

import (
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

var Instance = Config{}

func Read() (config Config) {
	config = readFromEnvVar()

	return
}

func readFromEnvVar() (config Config) {
	config.DigitalCurrency.ServerAddress = readEnvVarWithDefaultValue("DC_SERVER_ADDRESS", "http://localhost")
	config.DigitalCurrency.ApiKey = readEnvVarWithDefaultValue("DC_API_KEY", "apiKey")

	config.Auth.ServerAddress = readEnvVarWithDefaultValue("AUTH_SERVER_ADDRESS", "http://localhost")

	return
}

func readEnvVarWithDefaultValue(key, defaultValue string) string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue
	}
	return defaultValue
}

package presenter

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

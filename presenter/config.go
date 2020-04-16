package presenter

type (
	Config struct {
		DigitalCurrency DigitalCurrency `json:"digital_currency"`
	}

	DigitalCurrency struct {
		ServerAddress string `json:"server_address"`
		ApiKey        string `json:"api_key"`
	}
)

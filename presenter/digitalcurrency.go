package presenter

import (
	"context"
	"github.com/tech-showcase/financial-service/config"
	"github.com/tech-showcase/financial-service/controller"
	"github.com/tech-showcase/financial-service/model"
	digitalCurrencyProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
)

type digitalCurrencyServer struct{}

func NewDigitalCurrencyServer() digitalCurrencyProto.DigitalCurrencyServer {
	instance := digitalCurrencyServer{}

	return &instance
}

func (instance *digitalCurrencyServer) ConvertToSpecificCurrency(ctx context.Context, req *digitalCurrencyProto.ConvertCurrencyRequest) (resp *digitalCurrencyProto.ConvertCurrencyResponse, err error) {
	serverAddress := config.Instance.DigitalCurrency.ServerAddress
	apiKey := config.Instance.DigitalCurrency.ApiKey
	digitalCurrencyRepo := model.NewDigitalCurrencyRepo(serverAddress, apiKey)

	request := controller.ConvertRequest{
		Amount:        req.Amount,
		BaseCurrency:  req.BaseCurrency,
		QuoteCurrency: req.QuoteCurrency,
	}

	response, err := controller.ConvertToSpecificCurrency(request, digitalCurrencyRepo)
	if err != nil {
		return
	}

	resp = &digitalCurrencyProto.ConvertCurrencyResponse{}
	resp.Amount = response.Amount

	return
}

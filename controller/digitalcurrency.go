package controller

import (
	"context"
	"github.com/tech-showcase/financial-service/global"
	"github.com/tech-showcase/financial-service/model"
	dcProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
)

type digitalCurrencyServer struct{}

func NewDigitalCurrencyServer() dcProto.DigitalCurrencyServer {
	instance := digitalCurrencyServer{}

	return &instance
}

func (instance *digitalCurrencyServer) ConvertToSpecificCurrency(ctx context.Context, req *dcProto.ConvertToSpecificCurrencyRequest) (resp *dcProto.ConvertToSpecificCurrencyResponse, err error) {
	serverAddress := global.Configuration.DigitalCurrency.ServerAddress
	apiKey := global.Configuration.DigitalCurrency.ApiKey
	dcModel := model.NewDCBlueprint(serverAddress, apiKey)

	amount := req.Amount
	baseCurrency := req.BaseCurrency
	quoteCurrency := req.QuoteCurrency
	convertedAmount, err := dcModel.Convert(amount, baseCurrency, quoteCurrency)
	if err != nil {
		return
	}

	resp = &dcProto.ConvertToSpecificCurrencyResponse{}
	resp.Amount = float32(convertedAmount)
	resp.BaseCurrency = baseCurrency
	resp.QuoteCurrency = quoteCurrency

	return
}

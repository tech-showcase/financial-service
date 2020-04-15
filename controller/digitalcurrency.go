package controller

import (
	"context"
	"github.com/tech-showcase/financial-service/model"
	dcProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
)

type DigitalCurrencyServer struct{}

func NewDigitalCurrencyServer() dcProto.DigitalCurrencyServer {
	instance := DigitalCurrencyServer{}

	return &instance
}

func (instance *DigitalCurrencyServer) ConvertToSpecificCurrency(ctx context.Context, req *dcProto.ConvertToSpecificCurrencyRequest) (resp *dcProto.ConvertToSpecificCurrencyResponse, err error) {
	serverAddress := "http://rest-sandbox.coinapi.io/"
	apiKey := "A82B67D9-AF26-422D-87DE-8FED914E436E"
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

package controller

import (
	"github.com/tech-showcase/financial-service/model"
)

type (
	ConvertToSpecificCurrencyReq struct {
		Amount        float64
		BaseCurrency  string
		QuoteCurrency string
	}

	ConvertToSpecificCurrencyResp struct {
		Amount float64
	}
)

func ConvertToSpecificCurrency(req ConvertToSpecificCurrencyReq, digitalCurrencyRepo model.DigitalCurrencyRepo) (resp ConvertToSpecificCurrencyResp, err error) {
	convertedAmount, err := digitalCurrencyRepo.Convert(req.Amount, req.BaseCurrency, req.QuoteCurrency)
	if err != nil {
		return
	}

	resp.Amount = convertedAmount

	return
}

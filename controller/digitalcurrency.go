package controller

import (
	"github.com/tech-showcase/financial-service/model"
)

type (
	ConvertRequest struct {
		Amount        float64
		BaseCurrency  string
		QuoteCurrency string
	}

	ConvertResponse struct {
		Amount float64
	}
)

func ConvertToSpecificCurrency(req ConvertRequest, digitalCurrencyRepo model.DigitalCurrencyRepo) (resp ConvertResponse, err error) {
	convertedAmount, err := digitalCurrencyRepo.Convert(req.Amount, req.BaseCurrency, req.QuoteCurrency)
	if err != nil {
		return
	}

	resp.Amount = convertedAmount

	return
}

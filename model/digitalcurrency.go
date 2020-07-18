package model

import (
	"encoding/json"
	"fmt"
	"github.com/tech-showcase/financial-service/helper"
	"io/ioutil"
	"net/http"
)

type (
	digitalCurrencyRepo struct {
		ServerAddress string
		ApiKey        string
	}
	DigitalCurrencyRepo interface {
		Convert(value int64, assetIdBase, assetIdQuote string) (result float64, err error)
	}

	ResponseGetSpecificRate struct {
		AssetIDBase  string  `json:"asset_id_base"`
		AssetIDQuote string  `json:"asset_id_quote"`
		Rate         float64 `json:"rate"`
		Time         string  `json:"time"`
	}
)

func NewDigitalCurrencyRepoRepo(serverAddress, apiKey string) DigitalCurrencyRepo {
	instance := digitalCurrencyRepo{}
	instance.ServerAddress = serverAddress
	instance.ApiKey = apiKey

	return &instance
}

func (instance *digitalCurrencyRepo) Convert(value int64, assetIdBase, assetIdQuote string) (result float64, err error) {
	endpoint, _ := helper.JoinURL(
		instance.ServerAddress,
		fmt.Sprintf("/v1/exchangerate/%s/%s", assetIdBase, assetIdQuote))

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-CoinAPI-Key", instance.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseStruct := ResponseGetSpecificRate{}
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		return
	}

	result = float64(value) * responseStruct.Rate

	return
}

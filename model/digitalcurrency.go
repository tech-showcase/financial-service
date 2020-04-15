package model

import (
	"encoding/json"
	"github.com/tech-showcase/financial-service/helper"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	DCBlueprint struct {
		ServerAddress string
		ApiKey        string
	}
	DCInterface interface {
		Convert(value int64, base, quote string) (result float64, err error)
	}

	ResponseGetSpecificRate struct {
		AssetIDBase  string  `json:"asset_id_base"`
		AssetIDQuote string  `json:"asset_id_quote"`
		Rate         float64 `json:"rate"`
		Time         string  `json:"time"`
	}
)

func NewDCBlueprint(serverAddress, apiKey string) DCInterface {
	instance := DCBlueprint{}
	instance.ServerAddress = serverAddress
	instance.ApiKey = apiKey

	return &instance
}

func (instance *DCBlueprint) Convert(value int64, assetIdBase, assetIdQuote string) (result float64, err error) {
	endpoint, err := helper.JoinURL(
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

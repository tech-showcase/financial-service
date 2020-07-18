package model

import (
	"encoding/json"
	"fmt"
	"github.com/tech-showcase/financial-service/helper"
	"io/ioutil"
	"net/http"
)

type (
	dcRepo struct {
		ServerAddress string
		ApiKey        string
	}
	DCRepo interface {
		Convert(value int64, assetIdBase, assetIdQuote string) (result float64, err error)
	}

	ResponseGetSpecificRate struct {
		AssetIDBase  string  `json:"asset_id_base"`
		AssetIDQuote string  `json:"asset_id_quote"`
		Rate         float64 `json:"rate"`
		Time         string  `json:"time"`
	}
)

func NewDCRepo(serverAddress, apiKey string) DCRepo {
	instance := dcRepo{}
	instance.ServerAddress = serverAddress
	instance.ApiKey = apiKey

	return &instance
}

func (instance *dcRepo) Convert(value int64, assetIdBase, assetIdQuote string) (result float64, err error) {
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

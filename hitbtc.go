package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const btcurl string = "https://api.hitbtc.com/api/2/public/ticker"

//CryptoModel holds various crypto details
type CryptoModel struct {
	Id          string `json:"symbol"`
	FullName    string `json:"fullName"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

//GetCurrencyDetails gets currency details from API endpoint
func GetCurrencyDetails(symbol string) (*CryptoModel, error) {

	url := btcurl + "/" + symbol
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//unmarshall response if no error
	model := CryptoModel{}
	err := json.NewDecoder(resp.Body).Decode(&model)

	if err != nil {
		log.Printf("error in decoding the response %v", err)
	}
	return nil, err

	log.Printf("value of response is %v", model)
	return &model, nil

}

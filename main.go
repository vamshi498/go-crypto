package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"encoding/json"
	"github.com/gorilla/mux"
)

const btcurl string = "https://api.hitbtc.com/api/2/public/ticker"

func main() {
	fmt.Println("hello world")

	// channel for capturing os signals
	sChan := make(chan os.Signal, 1)

	//capture TERM and INT signals
	signal.Notify(sChan, syscall.SIGTERM)
	signal.Notify(sChan, syscall.SIGINT)

	r := mux.NewRouter()

	httpServer := &http.Server{
		Handler: r,
		Addr:    ":8081",
	}

	// Start httpServer
	go func() {
		log.Println("Starting Server")
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	r.HandleFunc("/currency/{symbol}", getCurrentPriceForSymbol).Methods(http.MethodGet)
	r.HandleFunc("/currency/all", getAllCurenciesForAllSymbols).Methods(http.MethodGet)
	//block untill we recieve signal
	<-sChan
	httpServer.Shutdown(context.Background())

	log.Println("Shutting down")
	os.Exit(0)

}

func getCurrentPriceForSymbol(w http.ResponseWriter, req *http.Request) {

	log.Print(req.URL.Path)
	vars := mux.Vars(req)

	//retrieve symbol from request
	symbol := vars["symbol"]
	log.Printf("value of symbol is %v", symbol)

	response, err := GetCurrencyDetails(symbol)

	if err != nil {
		log.Printf("error from getCurrentPriceForSymbol is %v", err)
	}
	// set response type as json
	w.Header().Set("Content-Type", "application/json")
	//converting the users slice to json
	json.NewEncoder(w).Encode(response)
}

func getAllCurenciesForAllSymbols(w http.ResponseWriter, req *http.Request) {

	
	response, err := GetAllCurrencies()

	if err != nil {
		log.Printf("error from getCurrentPriceForSymbol is %v", err)
	}
	// set response type as json
	w.Header().Set("Content-Type", "application/json")
	//converting the users slice to json
	json.NewEncoder(w).Encode(response)
}

//CryptoModel holds various crypto details
type CryptoModel struct {
	Symbol      string `json:"symbol"`
	FullName    string `json:"FullName"`
	Ask         string `json:"Ask"`
	Bid         string `json:"Bid"`
	Last        string `json:"Last"`
	Open        string `json:"Open"`
	Low         string `json:"Low"`
	High        string `json:"High"`
	FeeCurrency string `json:"FeeCurrency"`
}
//CryptoList
type CryptoList []CryptoModel
//GetCurrencyDetails gets currency details from API endpoint
func GetCurrencyDetails(symbol string) (*CryptoModel, error) {

	url := btcurl + "/" + symbol
	log.Printf("url is %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &CryptoModel{}
	err1 := json.NewDecoder(resp.Body).Decode(response)
	fmt.Println("response struct:", response)
	if err1 != nil {
		log.Printf("error in decoding the response %v", err)
		return nil, err1

	}

	return response, nil

}

//GetAllCurrencies gets all currencies 
func GetAllCurrencies() (*CryptoModel, error) {

	resp, err := http.Get("https://api.hitbtc.com/api/2/public/ticker")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	cryptoArr := &CryptoList{}
	log.Printf("response body is %v",resp.Body)
    err = json.NewDecoder(resp.Body).Decode(&cryptoArr)
	fmt.Printf("response structure:%+v", cryptoArr)
	if err != nil {
		log.Printf("error in decoding the response %v", err)
		return nil, err

	}

	return nil, nil

}

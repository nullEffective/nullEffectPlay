package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func channel() {

	messages := make(chan string, 2)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func convertFromUrl(amount float64) string {
	requestURL := "https://api.coinbase.com/v2/exchange-rates?currency=USD"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	var valueMap map[string]map[string]map[string]string //lose currency value
	json.Unmarshal(resBody, &valueMap)

	rates := valueMap["data"]["rates"]

	return convertFromMap(amount, rates)
}

func convertFromMap(amount float64, rates map[string]string) string {

	btcString := rates["BTC"]
	ethString := rates["ETH"]
	btc, _ := strconv.ParseFloat(btcString, 64)
	eth, _ := strconv.ParseFloat(ethString, 64)

	resultMap := make(map[string]interface{})
	resultMap["BTC"] = calc(amount, .7, btc)
	resultMap["ETH"] = calc(amount, .3, eth)
	resultMap["timestamp"] = time.Now()

	jsonBytes, jsonErr := json.MarshalIndent(resultMap, "", "   ")
	var json = "No json result"
	if jsonErr != nil {
		json = "Could not jsonize results"
	} else {
		json = string(jsonBytes)
	}
	return fmt.Sprintf("%v", json)
}

func calc(amount float64, percentage float64, coin float64) float64 {
	amt := amount * percentage * coin
	return amt
}

func main() {
	log.Println("----stephen leonard----")
	argsWithProg := os.Args
	if len(argsWithProg) == 0 {
		panic("Starting Dollar amount needed as program argument")
	}
	amount := os.Args[1]
	fmt.Println("Converting for: [$" + amount + "]")
	a, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(convertFromUrl(a))
}

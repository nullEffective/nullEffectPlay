package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func channel() {

	messages := make(chan string, 2)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func convert(amount float64) string {
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
	fmt.Println(req)
	return percentage(amount, resBody)
}

func percentage(amount float64, j []byte) string {

	var valueMap map[string]map[string]map[string]interface{}
	json.Unmarshal(j, &valueMap)

	rates := valueMap["data"]["rates"] //.(map[string]interface{})
	//fmt.Println("foo %v", rates)

	btcString := rates["BTC"].(string)
	ethString := rates["ETH"].(string)
	btc, _ := strconv.ParseFloat(btcString, 32)
	eth, _ := strconv.ParseFloat(ethString, 32)
	fmt.Printf("%v %v\n", btc, eth)

	/*
		b, err := json.MarshalIndent(valueMap, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}
		s := string(b)
		fmt.Print(s)
	*/
	btc70 := amount * 0.7
	eth30 := amount * 0.3
	resultBtc := btc70 * btc
	resultEth := eth30 * eth

	return fmt.Sprintf("foo %v %v", resultBtc, resultEth)
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
	fmt.Println(convert(a))
}

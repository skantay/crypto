package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Coin struct {
	Id string `json:"id"`
}

type CoinsResponse []Coin

func main() {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result CoinsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	for _, coin := range result {
		fmt.Println(coin.Id)
	}
}

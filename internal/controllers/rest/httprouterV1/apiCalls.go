package httprouterv1

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c controller) getCoinByApi() {
	url := "https://api.coingecko.com/api/v3/coins/list"
	//      https://marketdata.tradermade.com/api/v1/live_currencies_list?      api_key=     aiywJ90CwaNvhSMsYAvo
	req, err := http.Get(url)
	if err != nil {
		return
	}
	body, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
}

package httprouterv1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (c controller) rates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	w.Header().Set("Content-Type", "application/json")

	w.Write(body)
}

func (c controller) ratesCoin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	coin := p.ByName("coin")

	model, err := c.service.CoinService.GetCoin(r.Context(), strings.ToLower(coin))
	if err != nil {
		// if errors.Is(err, asd.ErrNoRecord) {
		// 	coins := []*asd.Coin{
		// 		&asd.Coin{

		// 		}
		// 	}

		// 	c.service.CoinService.CreateCoin(ctx, coins)
		// }
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", model)))
}

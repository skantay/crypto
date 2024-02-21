package httprouterv1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skantay/crypto/internal/domain/coin/model"
)

func (c controller) rates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return
}

func (c controller) ratesCoin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	coinToFind := p.ByName("coin")

	coin, err := c.service.CoinService.GetCoin(r.Context(), coinToFind)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			newCoin, err := getCoinData(coinToFind)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			errs := c.service.CoinService.CreateCoin(r.Context(), []model.Coin{*newCoin})
			if len(errs) != 0 {
				http.Error(w, errs[0].Error(), http.StatusInternalServerError)
				return
			}

			coin, err = c.service.CoinService.GetCoin(r.Context(), coinToFind)
		}
	}

	w.Write([]byte(fmt.Sprintf("%v", coin)))
}

package httprouterv1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skantay/crypto/internal/domain/coin/model"
)

type errorJSON struct {
	Error string `json:"error"`
}

func (c controller) rates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	coins, err := c.service.CoinService.GetMainCoins(r.Context())
	if err != nil {
		errorJSON := &errorJSON{
			Error: err.Error(),
		}

		js, err := json.Marshal(errorJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

		return
	}

	js, err := json.Marshal(map[string]any{"coins": &coins})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func (c controller) ratesCoin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	coinToFind := p.ByName("coin")

	coin, err := c.service.CoinService.GetCoin(r.Context(), coinToFind)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			newCoin, err := getCoinData(coinToFind)
			if err != nil {
				if errors.Is(err, model.ErrNoRecord) {
					errorJSON := &errorJSON{
						Error: err.Error(),
					}

					js, err := json.Marshal(errorJSON)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					w.Write(js)
					return
				}
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

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(&coin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

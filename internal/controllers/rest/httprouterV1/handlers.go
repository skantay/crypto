package httprouterv1

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skantay/crypto/internal/domain/coin/model"
)

func (c controller) rates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	coins, errs := c.service.CoinService.GetMainCoins(r.Context())
	if len(errs) != 0 {
		c.errorLog.Print(errs)
		if err := writeJSON(w, wrapJSON{"error": errs[0].Error()}); err != nil {
			c.errorLog.Print(err)
			internalServerError(w, err)
		}

		return
	}

	if err := writeJSON(w, wrapJSON{"coins": &coins}); err != nil {
		internalServerError(w, err)

		return
	}
}

func (c controller) ratesCoin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	coinToFind := p.ByName("coin")

	var coin model.Coin

	var err error

	coin, err = c.service.CoinService.GetCoin(r.Context(), coinToFind)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {

			// get by api
			coin, err = c.apiCalls.getCoin(r.Context(), coinToFind)
			// if error occured
			if err != nil {
				// if coin is not found by api then we respond with error
				if errors.Is(err, model.ErrNoRecord) {
					if err := writeJSON(w, wrapJSON{"error": err.Error()}); err != nil {
						internalServerError(w, err)
					}

					return
				}

				// if smt happens we respond with internal server error
				internalServerError(w, err)

				return
			}

			// create coin in db
			if coin.Price != 0.00 {
				if errs := c.service.CoinService.CreateCoin(r.Context(), []model.Coin{coin}); len(errs) != 0 && errs[0] != nil {
					internalServerError(w, errs[0])

					return
				}
			}
		} else {
			internalServerError(w, err)

			return
		}
	}

	if err := writeJSON(w, wrapJSON{"coin": coin}); err != nil {
		internalServerError(w, err)
	}
}

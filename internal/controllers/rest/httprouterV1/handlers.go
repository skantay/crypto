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
		// c.errorLog.Print(errs)
		// if err := wrapErrorJSON(w, errs); err != nil {
		// 	c.errorLog.Print(err)
		// 	internalServerError(w, err)
		// }

		// return
	}

	if err := writeJSON(w, wrapJSON{"coins": coins, "errors": errs}); err != nil {
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
			// to do a service which gets a coin, validates a coin and other stuff......
			return
		} else {
			internalServerError(w, err)

			return
		}
	}

	if err := writeJSON(w, wrapJSON{"coin": coin}); err != nil {
		internalServerError(w, err)
	}
}

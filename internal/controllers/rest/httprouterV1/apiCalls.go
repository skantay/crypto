package httprouterv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/skantay/crypto/internal/domain/coin/model"
)

type apiCalls interface {
	getCoin(ctx context.Context, name string) (model.Coin, error)
}

type coingecko struct{}

type CoinGeckoResponse struct {
	MarketData struct {
		CurrentPrice       map[string]float64 `json:"current_price"`
		PriceChangePercent map[string]float64 `json:"price_change_percentage_1h_in_currency"`
		Low                map[string]float64 `json:"low_24h"`
		High               map[string]float64 `json:"high_24h"`
	} `json:"market_data"`
	Error string `json:"error"`
}

func (c coingecko) getCoin(ctx context.Context, coinID string) (model.Coin, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", coinID)
	response, err := http.Get(url)
	if err != nil {
		return model.Coin{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return model.Coin{}, err
	}

	var coinData CoinGeckoResponse
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		return model.Coin{}, err
	}

	if coinData.Error != "" {
		return model.Coin{}, model.ErrNoRecord
	}

	result := model.Coin{
		Name:            coinID,
		Price:           coinData.MarketData.CurrentPrice["usd"],
		MinPrice:        coinData.MarketData.Low["usd"],
		MaxPrice:        coinData.MarketData.High["usd"],
		HourChangePrice: coinData.MarketData.PriceChangePercent["usd"],
	}

	return result, nil
}

//{"id":"ethereum","symbol":"eth","name":"Ethereum"}
//{"id":"bitcoin","symbol":"btc","name":"Bitcoin"}

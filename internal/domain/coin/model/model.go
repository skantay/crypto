package model

type Coin struct {
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	MinPrice        float64 `json:"min_price"`
	MaxPrice        float64 `json:"max_price"`
	HourChangePrice float64 `json:"hour_change_price"`
}

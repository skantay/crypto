package model

type Coin struct {
	Name            string  `json:"name"`
	Price           int     `json:"price"`
	MinPrice        int     `json:"min_price"`
	MaxPrice        int     `json:"max_price"`
	HourChangePrice float64 `json:"hour_change_price"`
}

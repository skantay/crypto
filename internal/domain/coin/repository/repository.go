package repository

import (
	"context"

	"github.com/skantay/crypto/internal/domain/coin/model"
)

type CoinRepository interface {
	CreateCoin(context.Context, model.Coin) error
	UpdateCoin(context.Context, model.Coin) error
	GetCoin(context.Context, string) (model.Coin, error)

	// to close db
	Close()
}

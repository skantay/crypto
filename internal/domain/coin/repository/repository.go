package repository

import (
	"context"

	"github.com/skantay/crypto/internal/domain/coin/model"
)

type CoinRepository interface {
	SaveCoin(context.Context, model.Coin) error
	UpdateCoin(context.Context, model.Coin) error
	GetCoin(context.Context, string) (model.Coin, error)
	GetMainCoins(context.Context, []string) ([]*model.Coin, error)

	// to close db
	Close()
}

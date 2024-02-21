package service

import (
	"context"

	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/repository"
)

type CoinService interface {
	CreateCoin(context.Context, []model.Coin) error
	UpdateCoin(context.Context, []model.Coin) error
	GetMainCoins(context.Context) ([]*model.Coin, error)
	GetCoin(context.Context, string) (model.Coin, error)
}

type coinService struct {
	repo repository.CoinRepository
}

func New(repo repository.CoinRepository) CoinService {
	return coinService{repo}
}

func (c coinService) CreateCoin(ctx context.Context, coins []model.Coin) error {
	for _, coin := range coins {
		if err := c.repo.SaveCoins(ctx, coin); err != nil {
			return
		}
	}
}

func (c coinService) UpdateCoin(context.Context, []model.Coin) error {
}

func (c coinService) GetMainCoins(context.Context) ([]*model.Coin, error) {
}

func (c coinService) GetCoin(context.Context, string) (model.Coin, error) {
}

package service

import (
	"context"
	"fmt"

	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/repository"
)

type CoinService interface {
	CreateCoin(context.Context, []model.Coin) []error
	UpdateCoin(context.Context, []model.Coin) []error
	GetMainCoins(context.Context) ([]*model.Coin, error)
	GetCoin(context.Context, string) (model.Coin, error)
}

type coinService struct {
	repo repository.CoinRepository
}

func New(repo repository.CoinRepository) CoinService {
	return coinService{repo}
}

func (c coinService) CreateCoin(ctx context.Context, coins []model.Coin) []error {
	var errs []error

	for _, coin := range coins {
		if err := c.repo.SaveCoin(ctx, coin); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (c coinService) UpdateCoin(ctx context.Context, coins []model.Coin) []error {
	var errs []error

	for _, coin := range coins {
		if err := c.repo.UpdateCoin(ctx, coin); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (c coinService) GetMainCoins(ctx context.Context) ([]*model.Coin, error) {
	coins := []string{
		"bitcoin",
		"ethereum",
	}

	result, err := c.repo.GetMainCoins(ctx, coins)
	if err != nil {
		return nil, fmt.Errorf("trouble with getting coins:%w", err)
	}

	return result, nil
}

func (c coinService) GetCoin(ctx context.Context, coin string) (model.Coin, error) {
	result, err := c.repo.GetCoin(ctx, coin)
	if err != nil {
		return model.Coin{}, fmt.Errorf("trouble with getting a coin:%w", err)
	}

	return result, nil
}

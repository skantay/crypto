package service

import (
	"context"
	"fmt"

	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/repository"
)

type CoinService interface {
	CreateCoin(ctx context.Context, coins []model.Coin) error
	UpdateCoin(ctx context.Context, coins []model.Coin) error
	GetMainCoins(ctx context.Context) ([]model.Coin, error)
	GetCoin(ctx context.Context, coin string) (model.Coin, error)
	GetAllCoins(ctx context.Context) ([]string, error)
}

type coinService struct {
	repo repository.CoinRepository
}

func New(repo repository.CoinRepository) CoinService {
	return coinService{repo}
}

func (c coinService) GetAllCoins(ctx context.Context) ([]string, error) {
	return c.repo.GetAllCoins(ctx)
}

func (c coinService) CreateCoin(ctx context.Context, coins []model.Coin) error {
	for _, coin := range coins {
		if _, err := c.repo.CreateCoin(ctx, coin); err != nil {
			return fmt.Errorf("%v: %w", coin.Name, err)
		}
	}

	return nil
}

func (c coinService) UpdateCoin(ctx context.Context, coins []model.Coin) error {
	for _, coin := range coins {
		if _, err := c.repo.UpdateCoin(ctx, coin); err != nil {
			return fmt.Errorf("%v: %w", coin.Name, err)
		}
	}

	return nil
}

func (c coinService) GetMainCoins(ctx context.Context) ([]model.Coin, error) {
	coins := []string{
		"bitcoin",
		"ethereum",
	}

	result := make([]model.Coin, 0, len(coins))

	for _, coin := range coins {
		gotCoin, err := c.repo.GetCoin(ctx, coin)
		if err != nil {
			return nil, fmt.Errorf("%v: %w", coin, err)
		}

		result = append(result, gotCoin)
	}

	return result, nil
}

func (c coinService) GetCoin(ctx context.Context, coin string) (model.Coin, error) {
	gotCoin, err := c.repo.GetCoin(ctx, coin)
	if err != nil {
		return model.Coin{}, fmt.Errorf("%v: %w", coin, err)
	}

	return gotCoin, nil
}

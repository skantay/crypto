package mock

import (
	"context"

	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/repository"
)

type mockUserRepository struct{}

func NewMockUserRepository() repository.CoinRepository {
	return &mockUserRepository{}
}

func (m mockUserRepository) CreateCoin(ctx context.Context, coin model.Coin) (model.Coin, error) {
	if coin.Name == "INVALID" {
		return coin, model.ErrNoPing
	}
	return coin, nil
}

func (m mockUserRepository) UpdateCoin(ctx context.Context, coin model.Coin) (model.Coin, error) {
	return coin, nil
}

func (m mockUserRepository) GetCoin(ctx context.Context, name string) (model.Coin, error) {
	return model.Coin{Name: name}, nil
}

func (m mockUserRepository) GetAllCoins(context.Context) ([]string, error) {
	return []string{"hello"}, nil
}

func (m mockUserRepository) Close() {
	return
}

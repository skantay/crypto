package service

import (
	"context"
	"errors"
	"testing"

	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/service/mock"
)

func TestCoinService_GetAllCoins(t *testing.T) {
	repo := mock.NewMockUserRepository()

	service := New(repo)

	coins, err := service.GetAllCoins(context.TODO())
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if coins[0] != "hello" {
		t.Error("unexxpected output")
	}
}

func TestCoinService_CreateCoin(t *testing.T) {
	repo := mock.NewMockUserRepository()

	service := New(repo)

	testCases := []struct {
		name        string
		coin        model.Coin
		expectedErr error
	}{
		{
			name: "positive create",
			coin: model.Coin{
				Name:            "BTC",
				Price:           0,
				MinPrice:        0,
				MaxPrice:        0,
				HourChangePrice: 0,
			},
			expectedErr: nil,
		},
		{
			name: "negative create",
			coin: model.Coin{
				Name:            "INVALID",
				Price:           0,
				MinPrice:        0,
				MaxPrice:        0,
				HourChangePrice: 0,
			},
			expectedErr: model.ErrNoPing,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CreateCoin(context.TODO(), []model.Coin{tt.coin})
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("\nexpected error %v\ngot error %v", tt.expectedErr, err)
			}
		})
	}
}

func TestCoinService_UpdateCoin(t *testing.T) {
	repo := mock.NewMockUserRepository()

	service := New(repo)

	testCases := []struct {
		name        string
		coin        model.Coin
		expectedErr error
	}{
		{
			name: "positive create",
			coin: model.Coin{
				Name:            "BTC",
				Price:           0,
				MinPrice:        0,
				MaxPrice:        0,
				HourChangePrice: 0,
			},
			expectedErr: nil,
		},
		{
			name: "negative create",
			coin: model.Coin{
				Name:            "INVALID",
				Price:           0,
				MinPrice:        0,
				MaxPrice:        0,
				HourChangePrice: 0,
			},
			expectedErr: model.ErrNoPing,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := service.UpdateCoin(context.TODO(), []model.Coin{tt.coin})
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("\nexpected error %v\ngot error %v", tt.expectedErr, err)
			}
		})
	}
}

func TestCoinService_GetMainCoins(t *testing.T) {
	repo := mock.NewMockUserRepository()

	service := New(repo)

	coins, err := service.GetMainCoins(context.TODO())
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if coins[0].Name != "bitcoin" && coins[1].Name != "ethereum" {
		t.Error("unexpected errors")
	}
}

func TestCoinService_GetCoin(t *testing.T) {
	repo := mock.NewMockUserRepository()

	service := New(repo)

	coins, err := service.GetCoin(context.TODO(), "BTC")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if coins.Name != "BTC" {
		t.Error("unexpected error")
	}
}

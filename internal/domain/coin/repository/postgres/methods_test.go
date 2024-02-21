package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/skantay/crypto/config"
	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/repository"
)

// Test SaveCoinsPositive function
// Positive Case
func TestSaveCoinsPositive(t *testing.T) {
	cfg := config.Database{
		Postgres: config.Postgres{
			User:     "user",
			Password: "pass",
			Host:     "localhost",
			DBName:   "domain_test",
			Port:     5432,
			SSLMode:  "disable",
		},
	}

	coinRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	coin := model.Coin{
		Name:            "BNB",
		Price:           150,
		MinPrice:        100,
		MaxPrice:        200,
		HourChangePrice: 1.5,
	}

	err = coinRepo.SaveCoins(ctx, coin)

	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		coinRepo.Close()
	})
}

// Test SaveCoinsPositive function
// Negative Case
func TestSaveCoinsNegative(t *testing.T) {
	cfg := config.Database{
		Postgres: config.Postgres{
			User:     "user",
			Password: "pass",
			Host:     "localhost",
			DBName:   "domain_test",
			Port:     5432,
			SSLMode:  "disable",
		},
	}

	coinRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	coin := model.Coin{}

	err = coinRepo.SaveCoins(ctx, coin)

	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		coinRepo.Close()
	})
}

// Test GetCoin function
// Positive Case
func TestGetCoinPositive(t *testing.T) {
	cfg := config.Database{
		Postgres: config.Postgres{
			User:     "user",
			Password: "pass",
			Host:     "localhost",
			DBName:   "domain_test",
			Port:     5432,
			SSLMode:  "disable",
		},
	}

	coinRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	coin := "BTC"

	exists := model.Coin{
		Name:            "BTC",
		Price:           10,
		MinPrice:        10,
		MaxPrice:        10,
		HourChangePrice: 10.5,
	}

	var result model.Coin

	result, err = coinRepo.GetCoin(ctx, coin)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	if result != exists {
		t.Errorf("wanted:%v\ngot:%v", exists, result)
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		coinRepo.Close()
	})
}

// Test GetCoin function
// Negative Case
func TestGetCoinNegative(t *testing.T) {
	cfg := config.Database{
		Postgres: config.Postgres{
			User:     "user",
			Password: "pass",
			Host:     "localhost",
			DBName:   "domain_test",
			Port:     5432,
			SSLMode:  "disable",
		},
	}

	coinRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	coin := ""

	_, err = coinRepo.GetCoin(ctx, coin)
	if err == nil {
		t.Error("Expected error")
	} else {
		if !errors.Is(err, model.ErrNoRecord) {
			t.Errorf("Expected ErrNoRecord, but got %v", err)
		}
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		coinRepo.Close()
	})
}

func TestGetMainCoins(t *testing.T) {
	cfg := config.Database{
		Postgres: config.Postgres{
			User:     "user",
			Password: "pass",
			Host:     "localhost",
			DBName:   "domain_test",
			Port:     5432,
			SSLMode:  "disable",
		},
	}

	coinRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	var result []*model.Coin

	result, err = coinRepo.GetMainCoins(ctx)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	if len(result) != 2 {
		t.Error("result seems wrong")
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		coinRepo.Close()
	})
}

func newTest(t *testing.T, cfg config.Database) (repository.CoinRepository, func(string, ...any) (sql.Result, error), error) {
	db, err := sql.Open("postgres", fmt.Sprintf(`
		user=%s
		password=%s
		dbname=%s
		host=%s
		port=%d
		sslmode=%s`,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
	))
	if err != nil {
		t.Fatal(err)
	}

	down, err := os.ReadFile("migrations/test.down.sql")
	if err != nil {
		t.Fatal(err)
	}

	up, err := os.ReadFile("migrations/test.up.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec(string(down)); err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec(string(up)); err != nil {
		t.Fatal(err)
	}

	return coinRepository{db}, db.Exec, nil
}

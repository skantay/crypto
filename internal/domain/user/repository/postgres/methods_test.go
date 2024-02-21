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
	"github.com/skantay/crypto/internal/domain/user/model"
	"github.com/skantay/crypto/internal/domain/user/repository"
)

func TestExistsUserPositive(t *testing.T) {
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

	userRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	ok, err := userRepo.ExistsUser(ctx, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !ok {
		t.Error("user not found")
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		userRepo.Close()
	})
}

func TestExistsUserNegative(t *testing.T) {
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

	userRepo, exec, err := newTest(t, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	ctx := context.Background()

	ok, err := userRepo.ExistsUser(ctx, 123)
	if err != nil {
		if !errors.Is(err, model.ErrNoRecord) {
			t.Errorf("Unexpected error: %v", err)
		}
		t.Logf("got error:%v", err)
	}

	if ok {
		t.Error("user found")
	}

	t.Cleanup(func() {
		down, err := os.ReadFile("migrations/test.down.sql")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := exec(string(down)); err != nil {
			t.Fatal(err)
		}
		userRepo.Close()
	})
}

func newTest(t *testing.T, cfg config.Database) (repository.UserRepository, func(string, ...any) (sql.Result, error), error) {
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

	return userRepository{db}, db.Exec, nil
}

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/skantay/crypto/config"
	"github.com/skantay/crypto/internal/domain/coin/model"
	"github.com/skantay/crypto/internal/domain/coin/repository"
)

type coinRepostiry struct {
	db *sql.DB
}

func (c coinRepostiry) Close() {
	c.db.Close()
}

func New(cfg config.Database) (repository.CoinRepository, error) {
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
		return nil, fmt.Errorf("postgres error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return coinRepostiry{db}, model.ErrNoPing
	}

	return coinRepostiry{db}, nil
}

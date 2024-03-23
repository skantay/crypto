package postgres

import (
	"database/sql"

	"github.com/skantay/crypto/internal/domain/coin/repository"
)

type coinRepository struct {
	db *sql.DB
}

func (c coinRepository) Close() {
	c.db.Close()
}

func New(db *sql.DB) repository.CoinRepository {
	return coinRepository{db}
}

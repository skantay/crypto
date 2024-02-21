package postgres

import (
	"database/sql"

	"github.com/skantay/crypto/internal/domain/user/repository"
)

type userRepository struct {
	db *sql.DB
}

func (c userRepository) Close() {
	c.db.Close()
}

func New(db *sql.DB) repository.UserRepository {
	return userRepository{db}
}

package repository

import (
	"context"

	"github.com/skantay/crypto/internal/domain/user/model"
)

type UserRepository interface {
	ExistsUser(context.Context, int64) (bool, error)
	CreateUser(context.Context, model.User) error
	UpdateUser(context.Context, model.User) error
	Delete(context.Context, int64, int64) error

	// to close DB
	Close()
}

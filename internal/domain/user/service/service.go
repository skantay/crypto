package service

import (
	"context"

	"github.com/skantay/crypto/internal/domain/user/repository"
)

type UserService interface {
	SetNotification(context.Context) error
	RemoveNotification(context.Context) error
}

type userService struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) UserService {
	return userService{repo}
}

func (u userService) SetNotification(context.Context) error {
	return nil
}

func (u userService) RemoveNotification(context.Context) error {
	return nil
}

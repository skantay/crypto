package domain

import (
	coinservice "github.com/skantay/crypto/internal/domain/coin/service"
	userservice "github.com/skantay/crypto/internal/domain/user/service"
)

type Service struct {
	coinService coinservice.CoinService
	userService userservice.UserService
}

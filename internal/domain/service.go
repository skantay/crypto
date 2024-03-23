package domain

import (
	"database/sql"

	coinrepo "github.com/skantay/crypto/internal/domain/coin/repository/postgres"
	coinservice "github.com/skantay/crypto/internal/domain/coin/service"
	userrepo "github.com/skantay/crypto/internal/domain/user/repository/postgres"
	userservice "github.com/skantay/crypto/internal/domain/user/service"
)

type Service struct {
	CoinService coinservice.CoinService
	UserService userservice.UserService
}

func New(db *sql.DB) Service {
	coinRepo := coinrepo.New(db)

	userRepo := userrepo.New(db)

	return Service{
		CoinService: coinservice.New(coinRepo),
		UserService: userservice.New(userRepo),
	}
}

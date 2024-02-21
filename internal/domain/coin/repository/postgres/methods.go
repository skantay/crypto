package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/skantay/crypto/internal/domain/coin/model"
)

func (c coinRepository) SaveCoins(ctx context.Context, coin model.Coin) error {
	stmt := `INSERT INTO coins(name, price, min_price, max_price, hour_change_price)
			VALUES($1, $2, $3, $4, $5);`

	if _, err := c.db.ExecContext(ctx, stmt,
		coin.Name,
		coin.Price,
		coin.MinPrice,
		coin.MaxPrice,
		coin.HourChangePrice); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	return nil
}

func (c coinRepository) GetCoin(ctx context.Context, coin string) (model.Coin, error) {
	coinResult := &model.Coin{}

	stmt := `SELECT * FROM coins WHERE name = $1`

	row := c.db.QueryRow(stmt, coin)

	if err := row.Scan(&coinResult.Name, &coinResult.Price, &coinResult.MinPrice, &coinResult.MaxPrice, &coinResult.HourChangePrice); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return *coinResult, model.ErrNoRecord
		}
		return *coinResult, err
	}

	return *coinResult, nil
}

func (c coinRepository) GetMainCoins(ctx context.Context) ([]*model.Coin, error) {
	result := []*model.Coin{}

	stmt := `SELECT * FROM coins WHERE name = 'BTC' OR name = 'ETH';`

	row, err := c.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		s := &model.Coin{}

		if err := row.Scan(&s.Name, &s.Price, &s.MinPrice, &s.MaxPrice, &s.HourChangePrice); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

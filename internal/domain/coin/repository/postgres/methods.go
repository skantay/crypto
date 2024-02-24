package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/skantay/crypto/internal/domain/coin/model"
)

func (c coinRepository) CreateCoin(ctx context.Context, coin model.Coin) error {
	stmt := `INSERT INTO domain.coins(name, price, min_price, max_price, hour_change_price)
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

func (c coinRepository) UpdateCoin(ctx context.Context, coin model.Coin) error {
	stmt := `UPDATE domain.coins SET price = $1,
				min_price = $2,
				max_price = $3,
				hour_change_price = $4
				WHERE name = $5;`

	if _, err := c.db.ExecContext(ctx, stmt,
		coin.Price,
		coin.MinPrice,
		coin.MaxPrice,
		coin.HourChangePrice,
		coin.Name); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	return nil
}

func (c coinRepository) GetCoin(ctx context.Context, coin string) (model.Coin, error) {
	coinResult := &model.Coin{}

	stmt := `SELECT * FROM domain.coins WHERE name = $1`

	row := c.db.QueryRowContext(ctx, stmt, coin)

	if err := row.Scan(
		&coinResult.Name,
		&coinResult.Price,
		&coinResult.MinPrice,
		&coinResult.MaxPrice,
		&coinResult.HourChangePrice,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return *coinResult, model.ErrNoRecord
		}

		return *coinResult, fmt.Errorf("query row error: %w", err)
	}

	return *coinResult, nil
}

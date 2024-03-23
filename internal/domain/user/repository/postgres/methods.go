package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/skantay/crypto/internal/domain/user/model"
)

func (u userRepository) ExistsUser(ctx context.Context, ID int64) (bool, error) {
	stmt := `SELECT id FROM users WHERE id = $1`

	row := u.db.QueryRowContext(ctx, stmt, ID)

	var id int64

	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, model.ErrNoRecord
		}

		return false, err
	}

	return ID == id, nil
}

func (u userRepository) CreateUser(ctx context.Context, user model.User) error {
	stmt := `INSERT INTO users(id, chat_id, coin, notification_interval, last_notification_time)
	VALUES ($1, $2, $3, $4, $5)`

	if _, err := u.db.ExecContext(ctx, stmt,
		user.ID,
		user.ChatID,
		user.Coin,
		user.NotificationInterval,
		user.LastNotificationTime,
	); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	return nil
}

func (u userRepository) UpdateUser(ctx context.Context, user model.User) error {
	stmt := `UPDATE users SET coin = $1, notification_interval = $2, last_notification_time = $3 WHERE id = $4 AND chat_id = $5`

	if _, err := u.db.ExecContext(ctx, stmt,
		user.Coin,
		user.NotificationInterval,
		user.LastNotificationTime,
		user.ID,
		user.ChatID,
	); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}
	return nil
}

func (u userRepository) Delete(ctx context.Context, id, chat_id int64) error {
	stmt := `DELETE FROM users WHERE id = $1 AND chat_id = $2`

	if _, err := u.db.ExecContext(ctx, stmt, id, chat_id); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}
	return nil
}

package model

import "time"

type User struct {
	ID                   int64         `json:"id"`
	ChatID               int64         `json:"chat_id"`
	Coin                 string        `json:"coin"`
	NotificationInterval time.Duration `json:"notification_interval"`
	LastNotificationTime time.Time     `json:"last_notification_time"`
}

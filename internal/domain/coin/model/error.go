package model

import "errors"

var (
	ErrNoPing   = errors.New("no ping")
	ErrNoRecord = errors.New("record not found")

	ErrEmptySlice = errors.New("empty slice")
)

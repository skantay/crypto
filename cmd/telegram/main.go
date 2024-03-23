package main

import "github.com/skantay/crypto/internal/apps/telegram"

func main() {
	if err := telegram.Run(); err != nil {
		panic(err)
	}
}

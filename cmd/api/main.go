package main

import "github.com/skantay/crypto/internal/apps/api"

func main() {
	if err := api.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	_ "net/http/pprof"

	"github.com/skantay/crypto/internal/apps/api"
)

func main() {
	if err := api.Run(); err != nil {
		panic(err)
	}
}

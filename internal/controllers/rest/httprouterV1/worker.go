package httprouterv1

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/skantay/crypto/internal/domain/coin/model"
)

func (c controller) refreshCoins(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			c.infoLog.Print("updating started")
			coins, _ := c.service.CoinService.GetAllCoins(ctx)
			wg.Add(len(coins))

			for _, coin := range coins {
				go func(coin string) {
					defer wg.Done()
					modelCoin, _ := c.apiCalls.getCoin(ctx, coin)
					c.infoLog.Printf("%s updated", coin)
					_ = c.service.CoinService.UpdateCoin(ctx, []model.Coin{modelCoin})
				}(coin)
			}

		case <-ctx.Done():
			fmt.Println("Context canceled, stopping refreshCoins")
			break
		}
	}

	wg.Wait()
}

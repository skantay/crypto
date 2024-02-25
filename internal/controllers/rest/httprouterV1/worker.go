package httprouterv1

import (
	"context"
	"errors"
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
					modelCoin, _, err := c.apiCalls.getCoin(ctx, coin)
					if err != nil {
						if errors.Is(err, model.ErrNoRecord) {
							c.errorLog.Printf("error during updating: %v", err)
							return
						}
						c.errorLog.Printf("error during updating: %v", err)
					}
					errs := c.service.CoinService.UpdateCoin(ctx, []model.Coin{modelCoin})
					if len(errs) != 0 && errs[0] != nil {
						c.errorLog.Printf("error during updating: %v", err)
					} else {
						c.infoLog.Printf("%s updated", coin)
					}
				}(coin)
			}

			wg.Wait()
			c.infoLog.Print("updating ended")
		case <-ctx.Done():
			fmt.Println("Context canceled, stopping refreshCoins")
		}
	}
}

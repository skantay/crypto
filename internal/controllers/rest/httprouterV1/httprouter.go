package httprouterv1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/skantay/crypto/config"
	"github.com/skantay/crypto/internal/controllers"
	"github.com/skantay/crypto/internal/domain"
)

type controller struct {
	service  domain.Service
	apiCalls apiCalls
	infoLog  *log.Logger
	errorLog *log.Logger
	cfg      config.Config
}

func New(
	service domain.Service,
	infoLog *log.Logger,
	errorLog *log.Logger,
	cfg config.Config,
) controllers.Controllers {
	return controller{
		service:  service,
		apiCalls: coingecko{},
		infoLog:  infoLog,
		errorLog: errorLog,
		cfg:      cfg,
	}
}

func (c controller) Run() error {
	router := httprouter.New()

	router.GET("/rates", c.rates)
	router.GET("/rates/:coin", c.ratesCoin)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", c.cfg.Server.Port),
		Handler: router,
	}

	go c.refreshCoins(context.Background(), time.Minute*1)

	c.infoLog.Print("strarting server on localhost:" + c.cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

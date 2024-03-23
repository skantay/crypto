package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/skantay/crypto/config"
	httprouterv1 "github.com/skantay/crypto/internal/controllers/rest/httprouterV1"
	"github.com/skantay/crypto/internal/domain"
)

func Run() error {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", fmt.Sprintf(`
		user=%s
		password=%s
		dbname=%s
		host=%s
		port=%d
		sslmode=%s`,
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.DBName,
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.SSLMode,
	))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	service := domain.New(db)

	infoLog := log.New(os.Stdout, "INFO:", log.Llongfile|log.Ldate)
	errorLog := log.New(os.Stdout, "ERROR:", log.Llongfile|log.Ldate)

	controller := httprouterv1.New(service, infoLog, errorLog, cfg)

	if err := controller.Run(); err != nil {
		return err
	}

	return nil
}

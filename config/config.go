package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
	API      API      `yaml:"api"`
}

type Database struct {
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type API struct {
	CoinMarket string `yaml:"coinMarket"`
	TraderMade string `yaml:"traderMade`
	Telegram   string `yaml:"telegram"`
}

func Load(path string) (Config, error) {
	viper.SetConfigFile(path)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	fmt.Println(config)
	return config, nil
}

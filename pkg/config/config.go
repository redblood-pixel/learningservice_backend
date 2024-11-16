package config

import (
	"log"

	"github.com/redblood-pixel/learning-service-go/internal/hash"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	postgres_repo "github.com/redblood-pixel/learning-service-go/pkg/repository/postgres"
	"github.com/redblood-pixel/learning-service-go/pkg/server"
	"github.com/spf13/viper"
)

type Cfg struct {
	HTTP server.Config
	DB   postgres_repo.Config
	Auth hash.Config
	JWT  tokenutil.Config
}

func NewCfg() *Cfg {
	var cfg Cfg
	var err error
	viper.SetConfigFile("configs/config.yaml")

	viper.SetDefault("server_address", "localhost:8000")

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if err = viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		log.Fatal(err.Error())
	}

	if err = viper.UnmarshalKey("auth.jwt", &cfg.JWT); err != nil {
		log.Fatal(err.Error())
	}

	if err = viper.UnmarshalKey("auth.password_salt", &cfg.Auth.PasswordSalt); err != nil {
		log.Fatal(err.Error())
	}

	if err = viper.UnmarshalKey("postgres", &cfg.DB); err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}

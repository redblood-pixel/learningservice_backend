package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type (
	Cfg struct {
		HTTP
		PostgresDB
		Auth
	}

	HTTP struct {
		ServerAdress   string        `mapstructure:"server_address"`
		ServerPort     string        `mapstructure:"server_port"`
		ContextTimeout time.Duration `mapstructure:"context_timeout"`
	}

	PostgresDB struct {
		DBHost     string `mapstructure:"db_host"`
		DBPort     string `mapstructure:"db_port"`
		DBName     string `mapstructure:"db_name"`
		DBUser     string `mapstructure:"db_user"`
		DBPassword string `mapstructure:"db_password"`
		DBSSLMode  string `mapstructure:"db_sslmode"`
	}

	Auth struct {
		JWT
		PasswordSalt string `mapstructure:"password_salt"`
	}

	JWT struct {
		AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expiry_time"`
		RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expiry_time"`
		SigningKey             string        `mapstructure:"signing_key"`
	}
)

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

	if err = viper.UnmarshalKey("auth.password_salt", &cfg.PasswordSalt); err != nil {
		log.Fatal(err.Error())
	}

	if err = viper.UnmarshalKey("postgres", &cfg.PostgresDB); err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}

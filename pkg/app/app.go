package app

import (
	"fmt"

	"github.com/redblood-pixel/learning-service-go/internal/hash"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/redblood-pixel/learning-service-go/pkg/config"
	"github.com/redblood-pixel/learning-service-go/pkg/handler"
	"github.com/redblood-pixel/learning-service-go/pkg/repository"
	gorm_repo "github.com/redblood-pixel/learning-service-go/pkg/repository/gorm"
	"github.com/redblood-pixel/learning-service-go/pkg/server"
	"github.com/redblood-pixel/learning-service-go/pkg/service"
	"github.com/sirupsen/logrus"
)

func Run() {

	cfg := config.NewCfg()

	fmt.Println(cfg.Auth.JWT.AccessTokenExpiryTime)

	postgres_db, err := gorm_repo.NewDB(&cfg.PostgresDB)
	if err != nil {
		logrus.Error(err)
		return
	}

	tokenManager := tokenutil.NewTokenManager(cfg.Auth.JWT.AccessTokenExpiryTime, cfg.Auth.JWT.RefreshTokenExpiryTime, cfg.Auth.JWT.SigningKey)

	hasher := hash.NewHasher(cfg.Auth.PasswordSalt)

	repos, err := repository.NewRepositories(postgres_db)
	if err != nil {
		logrus.Error(err)
		return
	}

	srvc := service.NewService(service.Dependencies{
		Repos:        repos,
		Hasher:       hasher,
		TokenManager: tokenManager,
	})

	handlers := handler.NewHandler(srvc, tokenManager)

	srv := server.NewServer(cfg, handlers.Init())

	srv.Run()
}

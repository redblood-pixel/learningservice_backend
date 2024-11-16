package app

import (
	"fmt"

	"github.com/redblood-pixel/learning-service-go/internal/hash"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/redblood-pixel/learning-service-go/pkg/config"
	"github.com/redblood-pixel/learning-service-go/pkg/handler"
	"github.com/redblood-pixel/learning-service-go/pkg/repository"
	postgres_repo "github.com/redblood-pixel/learning-service-go/pkg/repository/postgres"
	"github.com/redblood-pixel/learning-service-go/pkg/server"
	"github.com/redblood-pixel/learning-service-go/pkg/service"
	"github.com/sirupsen/logrus"
)

func Run() {

	cfg := config.NewCfg()

	postgres_db, err := postgres_repo.NewPostgresDB(&cfg.DB)
	if err != nil {
		logrus.Errorf("error occured while creating postgres_db: %v", err)
		return
	}
	fmt.Println(postgres_db)

	tokenManager := tokenutil.NewTokenManager(&cfg.JWT)

	hasher := hash.NewHasher(&cfg.Auth)

	repos, err := repository.NewRepository(postgres_db)
	if err != nil {
		logrus.Errorf("error occured while creating repostirory: %v", err)
		return
	}

	srvc := service.NewService(service.Dependencies{
		Repos:        repos,
		Hasher:       hasher,
		TokenManager: tokenManager,
	})

	handlers := handler.NewHandler(srvc, tokenManager)

	srv := server.NewServer(&cfg.HTTP, handlers.Init())

	srv.Run()
}

package server

import (
	"context"
	"net/http"
	"time"

	"github.com/redblood-pixel/learning-service-go/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Cfg, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.ServerPort,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (sv *Server) Run() {
	sv.httpServer.ListenAndServe()
}

func (sv *Server) Shutdown(ctx context.Context) error {
	return sv.httpServer.Shutdown(ctx)
}

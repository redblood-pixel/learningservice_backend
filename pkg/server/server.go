package server

import (
	"context"
	"net/http"
	"time"
)

type Config struct {
	ServerAdress   string        `mapstructure:"server_address"`
	ServerPort     string        `mapstructure:"server_port"`
	ContextTimeout time.Duration `mapstructure:"context_timeout"`
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           "0.0.0.0:" + cfg.ServerPort,
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

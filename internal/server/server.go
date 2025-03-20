package server

import (
	"context"
	"gitlab.com/petprojects9964409/todo_app/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handler http.Handler, cfg *config.Config) error {
	s.httpServer = &http.Server{
		Addr:           cfg.ServerHost + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

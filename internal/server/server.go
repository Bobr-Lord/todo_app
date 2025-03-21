package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/petprojects9964409/todo_app/internal/config"
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
	logrus.Infof("server listening on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

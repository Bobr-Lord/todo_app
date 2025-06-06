package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"gitlab.com/petprojects9964409/todo_app/internal/config"
	"gitlab.com/petprojects9964409/todo_app/internal/handler"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
	"gitlab.com/petprojects9964409/todo_app/internal/server"
	"gitlab.com/petprojects9964409/todo_app/internal/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.New(cfg.Postgres)
	if err != nil {
		logrus.Fatal(err)
	}
	_ = db

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer()
	go func() {
		if err := srv.Run(cfg.Port, handlers.InitRoutes(), cfg); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error starting http server: %v", err)
		}
	}()

	logrus.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error shutting down server: %v", err)
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error closing database: %v", err)
	}
	logrus.Info("Server stopped")
}

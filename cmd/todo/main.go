package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/petprojects9964409/todo_app/internal/config"
	"gitlab.com/petprojects9964409/todo_app/internal/handler"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
	"gitlab.com/petprojects9964409/todo_app/internal/server"
	"gitlab.com/petprojects9964409/todo_app/internal/service"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err.Error())
	}

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")

	db, err := repository.New(cfg.Postgres)
	if err != nil {
		logrus.Fatal(err)
	}
	_ = db

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer()
	if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error starting http server: %v", err)
	}
}

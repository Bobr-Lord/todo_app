package main

import (
	"gitlab.com/petprojects9964409/todo_app/internal/config"
	"gitlab.com/petprojects9964409/todo_app/internal/handler"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
	"gitlab.com/petprojects9964409/todo_app/internal/repository/postgres"
	"gitlab.com/petprojects9964409/todo_app/internal/server"
	"gitlab.com/petprojects9964409/todo_app/internal/service"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	_ = db

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer()
	if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error starting http server: %v", err)
	}
}

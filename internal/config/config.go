package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.com/petprojects9964409/todo_app/internal/repository/postgres"
)

type Config struct {
	Port     string          `yaml:"port"`
	Postgres postgres.Config `yaml:"postgres"`
}

func NewConfig() (*Config, error) {
	const op = "config.NewConfig"
	var cfg Config
	if err := cleanenv.ReadConfig("./configs/config.yml", &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &cfg, nil
}

package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
)

type Config struct {
	Port       string `env:"SERVER_PORT" envDefault:"8080"`
	ServerHost string `env:"SERVER_HOST" required:"true"`
	Salt       string `env:"SALT" required:"true"`
	SigningKey string `env:"SIGNING_KEY" required:"true"`
	Postgres   repository.Config
}

func NewConfig() (*Config, error) {
	const op = "config.NewConfig"
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &cfg, nil

}

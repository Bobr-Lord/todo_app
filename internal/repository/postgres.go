package repository

import (
	"fmt"

	"github.com/golang-migrate/migrate"

	// Импортируем пакет для инициализации драйвера PostgreSQL в migrate
	_ "github.com/golang-migrate/migrate/database/postgres"

	"time"

	// Импортируем пакет для инициализации драйвера migrate
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	todoListsTable  = "todo_lists"
	userListsTable  = "user_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"postgres"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Database string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	SSLMode  string `env:"POSTGRES_SSL" envDefault:"disable"`
	Password string `env:"DB_PASSWORD" required:"true"`
}

func New(cfg Config) (*sqlx.DB, error) {
	const op = "postgres.New"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s:  %w", op, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := migration(cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

func migration(cfg Config) error {
	const op = "postgres.NewMigration"

	m, err := migrate.New("file://db/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User,
			cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode))

	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %v", op, err)
	}
	return nil
}

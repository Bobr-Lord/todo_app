package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock_auth.go -package=mocks

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userID int, list models.TodoList) (int, error)
	GetAll(userID int) ([]models.TodoList, error)
	GetByID(userID, listID int) (models.TodoList, error)
	Delete(userID, listID int) error
	Update(userID int, listID int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item models.TodoItem) (int, error)
	GetAll(userID int, listID int) ([]models.TodoItem, error)
	GetByID(userID int, itemID int) (models.TodoItem, error)
	Delete(userID int, itemID int) error
	Update(userID int, itemID int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}

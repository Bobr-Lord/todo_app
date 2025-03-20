package service

import (
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list models.TodoList) (int, error)
	GetAll(userID int) ([]models.TodoList, error)
	GetByID(userID, listID int) (models.TodoList, error)
	Delete(userID, listID int) error
	Update(userID int, listID int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(userID int, listID int, item models.TodoItem) (int, error)
	GetAll(userID int, listID int) ([]models.TodoItem, error)
	GetByID(userID int, itemID int) (models.TodoItem, error)
	Delete(userID int, itemID int) error
	Update(userID int, itemID int, input models.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}

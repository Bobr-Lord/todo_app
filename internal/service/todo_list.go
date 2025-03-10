package service

import (
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

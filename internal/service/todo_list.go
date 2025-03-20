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

func (s *TodoListService) Create(userID int, list models.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]models.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (models.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}

func (s *TodoListService) Update(userID int, listID int, input models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, listID, input)
}

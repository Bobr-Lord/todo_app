package service

import (
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID int, listID int, item models.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID int, listID int) ([]models.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}

func (s *TodoItemService) GetByID(userID int, itemID int) (models.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}

func (s *TodoItemService) Delete(userID int, itemID int) error {
	return s.repo.Delete(userID, itemID)
}

func (s *TodoItemService) Update(userID int, itemID int, input models.UpdateItemInput) error {
	return s.repo.Update(userID, itemID, input)
}

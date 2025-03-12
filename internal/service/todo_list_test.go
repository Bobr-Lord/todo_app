package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository/mocks"
	"testing"
)

func TestTodoListService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	list := models.TodoList{
		Id:          1,
		Title:       "test",
		Description: "test",
	}
	mockRepo.EXPECT().Create(userId, list).Return(1, nil).Times(1)
	id, err := listService.Create(userId, list)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}
func TestTodoListService_Create_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	list := models.TodoList{
		Id:          1,
		Title:       "test",
		Description: "test",
	}
	mockRepo.EXPECT().Create(userId, list).Return(0, errors.New("error")).Times(1)
	id, err := listService.Create(userId, list)

	assert.Error(t, err)
	assert.Equal(t, 0, id)
}

func TestTodoListService_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	spTodoList := []models.TodoList{}
	spTodoList = append(spTodoList, models.TodoList{
		Id:          1,
		Title:       "test",
		Description: "test",
	})
	mockRepo.EXPECT().GetAll(userId).Return(spTodoList, nil).Times(1)
	result, err := listService.GetAll(userId)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestTodoListService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	spTodoList := []models.TodoList{}
	spTodoList = append(spTodoList, models.TodoList{
		Id:          1,
		Title:       "test",
		Description: "test",
	})
	mockRepo.EXPECT().GetAll(userId).Return(nil, errors.New("error")).Times(1)
	result, err := listService.GetAll(userId)
	assert.Error(t, err)
	assert.Empty(t, result)

}

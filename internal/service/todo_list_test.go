package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository/mocks"
)

func TestTodoListService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userID := 1
	list := models.TodoList{
		ID:          1,
		Title:       "test",
		Description: "test",
	}
	mockRepo.EXPECT().Create(userID, list).Return(1, nil).Times(1)
	id, err := listService.Create(userID, list)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}
func TestTodoListService_Create_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userID := 1
	list := models.TodoList{
		ID:          1,
		Title:       "test",
		Description: "test",
	}
	mockRepo.EXPECT().Create(userID, list).Return(0, errors.New("error")).Times(1)
	id, err := listService.Create(userID, list)

	assert.Error(t, err)
	assert.Equal(t, 0, id)
}

func TestTodoListService_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userID := 1
	spTodoList := []models.TodoList{}
	spTodoList = append(spTodoList, models.TodoList{
		ID:          1,
		Title:       "test",
		Description: "test",
	})
	mockRepo.EXPECT().GetAll(userID).Return(spTodoList, nil).Times(1)
	result, err := listService.GetAll(userID)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestTodoListService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userID := 1
	spTodoList := []models.TodoList{}
	spTodoList = append(spTodoList, models.TodoList{
		ID:          1,
		Title:       "test",
		Description: "test",
	})
	mockRepo.EXPECT().GetAll(userID).Return(nil, errors.New("error")).Times(1)
	result, err := listService.GetAll(userID)
	assert.Error(t, err)
	assert.Empty(t, result)

}

func TestTodoListService_GetById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)

	userID := 1
	listID := 1
	list := models.TodoList{
		ID:          1,
		Title:       "test",
		Description: "test",
	}
	mockRepo.EXPECT().GetByID(userID, listID).Return(list, nil).Times(1)
	result, err := listService.GetByID(userID, listID)
	assert.NoError(t, err)
	assert.Equal(t, list, result)

}

func TestTodoListService_GetById_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)

	userID := 1
	listID := 1

	mockRepo.EXPECT().GetByID(userID, listID).Return(models.TodoList{}, errors.New("error")).Times(1)
	result, err := listService.GetByID(userID, listID)
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestTodoListService_Delete_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	listId := 1
	mockRepo.EXPECT().Delete(userId, listId).Return(nil).Times(1)
	err := listService.Delete(userId, listId)
	assert.NoError(t, err)
}

func TestTodoListService_Delete_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userID := 1
	listID := 1
	mockRepo.EXPECT().Delete(userID, listID).Return(errors.New("error")).Times(1)
	err := listService.Delete(userID, listID)
	assert.Error(t, err)
}

func TestTodoListService_Update_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	listId := 1
	title := "title"
	description := "description"
	input := models.UpdateListInput{
		Title:       &title,
		Description: &description,
	}
	mockRepo.EXPECT().Update(userId, listId, input).Return(nil).Times(1)
	err := listService.Update(userId, listId, input)
	assert.NoError(t, err)
}

func TestTodoListService_Update_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTodoList(ctrl)
	listService := NewTodoListService(mockRepo)
	userId := 1
	listId := 1
	title := "test"
	description := "test"
	input := models.UpdateListInput{
		Title:       &title,
		Description: &description,
	}
	mockRepo.EXPECT().Update(userId, listId, input).Return(errors.New("error"))
	err := listService.Update(userId, listId, input)
	assert.Error(t, err)
}

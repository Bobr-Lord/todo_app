package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository/mocks"
)

func TestAuthService_CreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthorization(ctrl)

	userService := NewAuthService(mockAuth)

	userServ := models.User{
		ID:       1,
		Name:     "test",
		Username: "testuser",
		Password: "testpassword",
	}
	userRepo := models.User{
		ID:       1,
		Name:     "test",
		Username: "testuser",
		Password: generatePasswordHash("testpassword"),
	}
	mockAuth.EXPECT().CreateUser(userRepo).Return(1, nil).Times(1)

	id, err := userService.CreateUser(userServ)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestAuthService_CreateUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthorization(ctrl)

	userService := NewAuthService(mockAuth)

	userServ := models.User{
		ID:       1,
		Name:     "test",
		Username: "testuser",
		Password: "testpassword",
	}
	userRepo := models.User{
		ID:       1,
		Name:     "test",
		Username: "testuser",
		Password: generatePasswordHash("testpassword"),
	}
	mockAuth.EXPECT().CreateUser(userRepo).Return(0, errors.New("error")).Times(1)
	id, err := userService.CreateUser(userServ)
	assert.Error(t, err)
	assert.Equal(t, "error", err.Error())
	assert.Equal(t, 0, id)
}

func TestAuthService_GenerateToken_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := mocks.NewMockAuthorization(ctrl)
	userService := NewAuthService(mockAuth)
	username := "testuser"
	password := "testpassword"

	user := models.User{
		ID:       1,
		Name:     username,
		Username: "testuser",
		Password: password,
	}

	mockAuth.EXPECT().GetUser(username, generatePasswordHash(password)).Return(user, nil).Times(1)
	token, err := userService.GenerateToken(username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	t.Logf("token: %s", token)
}

func TestAuthService_GenerateToken_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuth := mocks.NewMockAuthorization(ctrl)
	userService := NewAuthService(mockAuth)
	username := "testuser"
	password := "testpassword"

	mockAuth.EXPECT().GetUser(username, generatePasswordHash(password)).Return(models.User{}, errors.New("error generate")).Times(1)
	token, err := userService.GenerateToken(username, password)

	assert.Error(t, err)
	assert.Empty(t, token)
	t.Logf("error: %v", err)
}

package service

import (
	"crypto/sha1"
	"fmt"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"gitlab.com/petprojects9964409/todo_app/internal/repository"
)

const salt = "kasjgfeiuysaocjna,biljkjabdm,czmlwIU"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

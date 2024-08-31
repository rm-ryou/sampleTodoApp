package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type AuthServicer interface {
	SignIn(email, password string, isAdminResource bool) (*entity.UserResponse, error)
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) SignIn(email, password string, isAdminResource bool) (*entity.UserResponse, error) {
	return nil, nil
}

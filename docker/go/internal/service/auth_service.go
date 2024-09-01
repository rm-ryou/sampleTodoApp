package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type AuthServicer interface {
	SignUp(name, email, password string) (*entity.Auth, error)
	SignIn(email, password string, isAdminResource bool) (*entity.Auth, error)
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) SignUp(name, email, password string) (*entity.Auth, error) {
	return nil, nil
}

func (as *AuthService) SignIn(email, password string, isAdminResource bool) (*entity.Auth, error) {
	return nil, nil
}

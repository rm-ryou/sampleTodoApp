package service

import (
	"errors"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	"github.com/rm-ryou/sampleTodoApp/internal/storage/mysql/repository"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
	"github.com/rm-ryou/sampleTodoApp/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthServicer interface {
	SignUp(name, email, password string) (*entity.Auth, error)
	SignIn(email, password string, isAdminResource bool) (*entity.Auth, error)
}

type AuthService struct {
	r repository.IUserRepository
}

func NewAuthService(r repository.IUserRepository) *AuthService {
	return &AuthService{r}
}

func attachToken(user *entity.User) (*entity.Auth, error) {
	token, err := auth.GenerateToken(user.ID, utils.RealTime{})
	if err != nil {
		return nil, err
	}

	return &entity.Auth{User: *user, Accesstoken: token}, nil
}

func (as *AuthService) SignUp(name, email, password string) (*entity.Auth, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("invalid params")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user, err := as.r.CreateUser(&entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return attachToken(user)
}

func (as *AuthService) SignIn(email, password string, isAdminResource bool) (*entity.Auth, error) {
	return nil, nil
}

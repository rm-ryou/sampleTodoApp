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

func setToken(user *entity.User) (*entity.Auth, error) {
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

	hashed, err := EncryptPassword(password)
	if err != nil {
		return nil, err
	}

	if err := as.r.CreateUser(&entity.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}); err != nil {
		return nil, err
	}

	user, err := as.r.ReadUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return setToken(user)
}

func (as *AuthService) SignIn(email, password string, isAdminResource bool) (*entity.Auth, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid params")
	}

	user, err := as.r.ReadUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Admin != isAdminResource {
		return nil, errors.New("failed to sign in")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return setToken(user)
}

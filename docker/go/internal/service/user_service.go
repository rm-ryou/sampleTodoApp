package service

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	"github.com/rm-ryou/sampleTodoApp/internal/storage/mysql/repository"
)

type UserServicer interface {
	GetUser(id int) (*entity.UserResponse, error)
	GetUsers() ([]entity.UserResponse, error)
	EditUser(id int, name, email, password string) (*entity.UserResponse, error)
	DeleteUser(id int) error
}

type UserService struct {
	r repository.IUserRepository
}

func NewUserService(r repository.IUserRepository) *UserService {
	return &UserService{r}
}

func (us *UserService) GetUser(id int) (*entity.UserResponse, error) {
	return nil, nil
}

func (us *UserService) GetUsers() ([]entity.UserResponse, error) {
	return nil, nil
}

func (us *UserService) EditUser(id int, name, email, password string) (*entity.UserResponse, error) {
	return nil, nil
}

func (us *UserService) DeleteUser(id int) error {
	return nil
}

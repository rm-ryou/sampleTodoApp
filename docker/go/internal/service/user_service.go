package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type UserServicer interface {
	CreateUser(name, email, password string) (*entity.UserResponse, error)
	GetUser(id int) (*entity.UserResponse, error)
	GetUsers() ([]entity.UserResponse, error)
	EditUser(id int, name, email, password string) (*entity.UserResponse, error)
	DeleteUser(id int) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(name, email, password string) (*entity.UserResponse, error) {
	return nil, nil
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

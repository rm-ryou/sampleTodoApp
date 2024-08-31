package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type UserServicer interface {
	CreateUser(name, email, password string) (*entity.User, error)
	GetUser(id int) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	EditUser(id int, name, email, password string) (*entity.User, error)
	DeleteUser(id int) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	return nil, nil
}

func (us *UserService) GetUser(id int) (*entity.User, error) {
	return nil, nil
}

func (us *UserService) GetUsers() ([]entity.User, error) {
	return nil, nil
}

func (us *UserService) EditUser(id int, name, email, password string) (*entity.User, error) {
	return nil, nil
}

func (us *UserService) DeleteUser(id int) error {
	return nil
}

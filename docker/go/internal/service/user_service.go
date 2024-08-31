package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type UserServicer interface {
	CreateUser(name, email, password string) *entity.User
	GetUser(id int) *entity.User
	GetUsers() []entity.User
	EditUser(id int, name, email, password string) *entity.User
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(name, email, password string) *entity.User {
	return nil
}

func (us *UserService) GetUser(id int) *entity.User {
	return nil
}

func (us *UserService) GetUsers() []entity.User {
	return nil
}

func (us *UserService) EditUser(id int, name, email, password string) *entity.User {
	return nil
}

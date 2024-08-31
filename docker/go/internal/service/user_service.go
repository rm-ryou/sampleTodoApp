package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type UserServicer interface {
	GetUser(id int) *entity.User
	GetUsers() []entity.User
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUser(id int) *entity.User {
	return nil
}

func (us *UserService) GetUsers() []entity.User {
	return nil
}

package service

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type UserServicer interface {
	GetUsers() []entity.User
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUsers() []entity.User {
	return nil
}

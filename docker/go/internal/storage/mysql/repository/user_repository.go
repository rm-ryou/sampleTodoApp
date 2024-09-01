package repository

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}

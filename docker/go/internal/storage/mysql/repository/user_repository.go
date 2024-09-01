package repository

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

type IUserRepository interface {
	CreateUser(user *entity.User) error
	ReadUser(id int) (*entity.User, error)
	ReadUserByEmail(email string) (*entity.User, error)
	ReadUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id int) error
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(user *entity.User) error {
	return nil
}

func (ur *UserRepository) ReadUser(id int) (*entity.User, error) {
	return nil, nil
}

func (ur *UserRepository) ReadUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}

func (ur *UserRepository) ReadUsers() ([]entity.User, error) {
	return nil, nil
}

func (ur *UserRepository) UpdateUser(user *entity.User) error {
	return nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	return nil
}

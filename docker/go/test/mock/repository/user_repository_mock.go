package repository_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type userRepositoryMock struct{}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}

func (urm *userRepositoryMock) CreateUser(user *entity.User) (*entity.User, error) {
	tmpHashedPassword := user.Password

	savedUser, err := testdata.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	savedUser.Password = tmpHashedPassword

	return savedUser, nil
}

package repository_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"golang.org/x/crypto/bcrypt"
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

func (urm *userRepositoryMock) GetUserByEmail(email string) (*entity.User, error) {
	user, err := testdata.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return user, nil
}

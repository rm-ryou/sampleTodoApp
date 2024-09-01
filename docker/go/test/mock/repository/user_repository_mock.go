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

func (urm *userRepositoryMock) CreateUser(user *entity.User) error {
	tmpHashedPassword := user.Password

	savedUser, err := testdata.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	savedUser.Password = tmpHashedPassword

	return nil
}

func (urm *userRepositoryMock) ReadUser(id int) (*entity.User, error) {
	user := testdata.UserTestData[id-1]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return &user, nil
}

func (urm *userRepositoryMock) ReadUserByEmail(email string) (*entity.User, error) {
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

func (urm *userRepositoryMock) ReadUsers() ([]entity.User, error) {
	return testdata.UserTestData[1:], nil
}

func (urm *userRepositoryMock) UpdateUser(user *entity.User) error {
	return nil
}

func (urm *userRepositoryMock) DeleteUser(id int) error {
	return nil
}

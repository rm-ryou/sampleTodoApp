package service_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type userServiceMock struct{}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (usm *userServiceMock) CreateUser(name, email, password string) (*entity.UserResponse, error) {
	return &testdata.UserResponseTestData[1], nil
}

func (usm *userServiceMock) GetUser(id int) (*entity.UserResponse, error) {
	return &testdata.UserResponseTestData[1], nil
}

func (usm *userServiceMock) GetUsers() ([]entity.UserResponse, error) {
	return testdata.UserResponseTestData[1:], nil
}

func (usm *userServiceMock) EditUser(id int, name, email, password string) (*entity.UserResponse, error) {
	baseUser := testdata.UserResponseTestData[1]

	if name != "" {
		baseUser.Name = name
	}
	if email != "" {
		baseUser.Email = email
	}

	return &baseUser, nil
}

func (usm *userServiceMock) DeleteUser(id int) error {
	return nil
}

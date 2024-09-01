package service_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type userServiceMock struct{}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (usm *userServiceMock) GetUser(id int) (*entity.User, error) {
	return &testdata.UserTestData[1], nil
}

func (usm *userServiceMock) GetUsers() ([]entity.User, error) {
	return testdata.UserTestData[1:], nil
}

func (usm *userServiceMock) EditUser(id int, password, name, email, Newpassword string) (*entity.User, error) {
	baseUser := testdata.UserTestData[1]

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

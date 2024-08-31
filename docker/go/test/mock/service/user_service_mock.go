package service_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type userServiceMock struct{}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (usm *userServiceMock) CreateUser(name, email, password string) *entity.User {
	return &testdata.UserTestData[1]
}

func (usm *userServiceMock) GetUser(id int) *entity.User {
	return &testdata.UserTestData[1]
}

func (usm *userServiceMock) GetUsers() []entity.User {
	return testdata.UserTestData[1:]
}

func (usm *userServiceMock) EditUser(id int, name, email, password string) *entity.User {
	baseUser := testdata.UserTestData[1]

	if name != "" {
		baseUser.Name = name
	}
	if email != "" {
		baseUser.Email = email
	}
	if password != "" {
		baseUser.Password = password
	}

	return &baseUser
}

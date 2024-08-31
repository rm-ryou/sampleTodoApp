package service_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type userServiceMock struct{}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (usm *userServiceMock) GetUser(id int) *entity.User {
	return &testdata.UserTestData[1]
}

func (usm *userServiceMock) GetUsers() []entity.User {
	return testdata.UserTestData[1:]
}

func (usm *userServiceMock) EditUser(id int, name, email, password string) *entity.User {
	baseUser := testdata.UserTestData[1]

	if baseUser.Name != name {
		baseUser.Name = name
	}
	if baseUser.Email != email {
		baseUser.Email = email
	}
	if baseUser.Password != password {
		baseUser.Password = password
	}

	return &baseUser
}

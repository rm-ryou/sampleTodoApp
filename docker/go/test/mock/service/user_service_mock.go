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

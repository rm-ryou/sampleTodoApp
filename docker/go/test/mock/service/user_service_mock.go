package service_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type userServiceMock struct{}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (usm *userServiceMock) GetUsers() []entity.User {
	return testdata.UserTestData[1:]
}

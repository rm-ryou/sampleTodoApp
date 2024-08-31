package service_mock

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type authServiceMock struct {}

func NewAuthServiceMock() *authServiceMock {
	return &authServiceMock{}
}

func (asm *authServiceMock) SignIn(email, password string) (*entity.UserResponse, error) {
	return &testdata.UserResponseTestData[1], nil
}

package service_mock

import (
	"errors"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type authServiceMock struct{}

func NewAuthServiceMock() *authServiceMock {
	return &authServiceMock{}
}

func (asm *authServiceMock) SignIn(email, password string, isAdminResource bool) (*entity.UserResponse, error) {
	user, err := getUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Admin != isAdminResource {
		return nil, errors.New("failed to sign in")
	}
	return &testdata.UserResponseTestData[user.ID - 1], nil
}

func getUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	switch email {
	case "admin@example.com":
		user = testdata.UserTestData[0]
	case "user01@example.com":
		user = testdata.UserTestData[1]
	default:
		return nil, errors.New("user not found")
	}
	return &user, nil
}

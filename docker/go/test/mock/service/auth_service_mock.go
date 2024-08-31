package service_mock

import (
	"errors"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
	"github.com/rm-ryou/sampleTodoApp/pkg/utils"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
)

type authServiceMock struct{}

func NewAuthServiceMock() *authServiceMock {
	return &authServiceMock{}
}

func (asm *authServiceMock) SignUp(name, email, password string) (*entity.AuthResponse, error) {
	user := testdata.UserResponseTestData[1]
	token, _ := auth.GenerateToken(user.ID, utils.RealTime{})

	authResponse := &entity.AuthResponse{
		UserResponse: user,
		Accesstoken:  token,
	}
	return authResponse, nil
}

func (asm *authServiceMock) SignIn(email, password string, isAdminResource bool) (*entity.AuthResponse, error) {
	user, err := getUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Admin != isAdminResource {
		return nil, errors.New("failed to sign in")
	}

	userResponse := &testdata.UserResponseTestData[user.ID-1]
	token, _ := auth.GenerateToken(userResponse.ID, utils.RealTime{})

	authResponse := &entity.AuthResponse{
		UserResponse: *userResponse,
		Accesstoken:  token,
	}
	return authResponse, nil
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

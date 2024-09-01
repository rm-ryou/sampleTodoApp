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

func (asm *authServiceMock) SignUp(name, email, password string) (*entity.Auth, error) {
	user := testdata.UserTestData[1]
	token, _ := auth.GenerateToken(user.ID, utils.RealTime{})

	auth := &entity.Auth{
		User: user,
		Accesstoken:  token,
	}
	return auth, nil
}

func (asm *authServiceMock) SignIn(email, password string, isAdminResource bool) (*entity.Auth, error) {
	user, err := testdata.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Admin != isAdminResource {
		return nil, errors.New("failed to sign in")
	}

	token, _ := auth.GenerateToken(user.ID, utils.RealTime{})

	auth := &entity.Auth{
		User: *user,
		Accesstoken:  token,
	}
	return auth, nil
}

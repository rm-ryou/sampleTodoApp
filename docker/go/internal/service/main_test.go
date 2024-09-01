package service

import (
	"testing"

	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
	repository_mock "github.com/rm-ryou/sampleTodoApp/test/mock/repository"
)

var (
	as         *AuthService
	us         *UserService
	signingKey string
)

func TestMain(m *testing.M) {
	signingKey = "test"
	auth.InitializeSigningKey(signingKey)

	ur := repository_mock.NewUserRepositoryMock()
	as = NewAuthService(ur)
	us = NewUserService(ur)

	m.Run()
}

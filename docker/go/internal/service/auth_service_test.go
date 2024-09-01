package service

import (
	"testing"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUp(t *testing.T) {
	user := testdata.UserTestData[1]
	res, err := as.SignUp(user.Name, user.Email, user.Password)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, res.ID)
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password)))
	assert.NotEmpty(t, res.Accesstoken)
}

func TestSignIn(t *testing.T) {
	correctCases := []struct {
		title           string
		isAdminResource bool
		user            entity.User
	}{
		{
			title:           "user can sign in",
			isAdminResource: false,
			user:            testdata.UserTestData[1],
		},
		{
			title:           "admin can sign in",
			isAdminResource: true,
			user:            testdata.UserTestData[0],
		},
	}

	for _, test := range correctCases {
		t.Run(test.title, func(t *testing.T) {
			res, err := as.SignIn(test.user.Email, test.user.Password, test.isAdminResource)

			assert.Nil(t, err)
			assert.Equal(t, test.user.ID, res.ID)
			assert.NotEmpty(t, res.Accesstoken)
		})
	}

	inCorrectCases := []struct {
		title           string
		isAdminResource bool
		user            entity.User
	}{
		{
			title:           "user can't sign in as admin",
			isAdminResource: true,
			user:            testdata.UserTestData[1],
		},
		{
			title:           "admin can't sign in as user",
			isAdminResource: false,
			user:            testdata.UserTestData[0],
		},
		{
			title:           "user not exist",
			isAdminResource: false,
			user:            entity.User{Email: "hoge@example.com", Password: "hoge"},
		},
	}

	for _, test := range inCorrectCases {
		t.Run(test.title, func(t *testing.T) {
			res, err := as.SignIn(test.user.Email, test.user.Password, test.isAdminResource)

			assert.Nil(t, res)
			assert.NotNil(t, err)
		})
	}
}

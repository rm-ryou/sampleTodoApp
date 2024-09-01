package service

import (
	"testing"

	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUp(t *testing.T) {
	user := testdata.UserTestData[1]
	res, err := as.SignUp(user.Name, user.Email, user.Password)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, user.ID, res.ID)
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password)))
	assert.NotEmpty(t, res.Accesstoken)
}

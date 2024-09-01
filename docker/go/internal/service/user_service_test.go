package service

import (
	"testing"

	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
)

func TestEditUser(t *testing.T) {
	t.Run("valid password", func(t *testing.T) {
		user := testdata.UserTestData[1]
		res, err := us.EditUser(user.ID, user.Password, "NewName", "NewEmail", "NewPassword")

		assert.NotNil(t, res)
		assert.Nil(t, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		user := testdata.UserTestData[1]
		res, err := us.EditUser(user.ID,
			"InvalidPassword",
			"NewName",
			"NewEmail",
			"NewPassword")

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})
}

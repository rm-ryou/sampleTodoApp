package repository

import (
	"testing"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := entity.User{
		Name:     "testCreateUser",
		Email:    "test_create_user@example.com",
		Password: "testCreateUserPassword",
	}

	err := ur.CreateUser(&user)
	if err != nil {
		t.Error(err)
	}

	var count int
	err = testDB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(testdata.UserTestData)+1, count)

	t.Cleanup(func() {
		query := `
			DELETE FROM users
			WHERE
				name = ? AND
				email = ?
		`

		testDB.Exec(query, user.Name, user.Email)
	})
}

func TestReadUser(t *testing.T) {
	t.Run("valid user id", func(t *testing.T) {
		expectedUser := testdata.UserTestData[0]
		user, err := ur.ReadUser(expectedUser.ID)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectedUser.Name, user.Name)
	})

	t.Run("not exists user id", func(t *testing.T) {
		user, err := ur.ReadUser(0)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}

func TestReadUserByEmail(t *testing.T) {
	t.Run("valid user email", func(t *testing.T) {
		expectedUser := testdata.UserTestData[0]
		user, err := ur.ReadUserByEmail(expectedUser.Email)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectedUser.Name, user.Name)
	})

	t.Run("not exists user", func(t *testing.T) {
		user, err := ur.ReadUserByEmail("not_exist@example.com")

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}

func TestReadUsers(t *testing.T) {
	users, err := ur.ReadUsers()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(testdata.UserTestData), len(users))
}

func TestUpdateUser(t *testing.T) {
	user := testdata.UserTestData[0]
	newName := "newName"
	user.Name = newName

	err := ur.UpdateUser(&user)
	if err != nil {
		t.Error(err)
	}

	var savedUserName string
	err = testDB.QueryRow("SELECT name FROM users WHERE ID = ?", user.ID).Scan(&savedUserName)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, newName, savedUserName)
}

func TestDeleteUser(t *testing.T) {
	deleteUser := testdata.UserTestData[0]
	err := ur.DeleteUser(deleteUser.ID)
	if err != nil {
		t.Error(err)
	}

	var count int
	err = testDB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(testdata.UserTestData)-1, count)

	t.Cleanup(func() {
		query := `
			INSERT
			INTO users
				(name, email, password)
			VALUES
				(?, ?, ?)
		`
		stmt, _ := testDB.Prepare(query)
		defer stmt.Close()
		stmt.Exec(deleteUser.Name, deleteUser.Email, deleteUser.Password)
	})
}

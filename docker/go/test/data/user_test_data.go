package testdata

import (
	"errors"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
)

var UserTestData = []entity.User{
	{
		ID:       1,
		Name:     "admin",
		Email:    "admin@example.com",
		Password: "admin",
		Admin:    true,
	},
	{
		ID:       2,
		Name:     "user01",
		Email:    "user01@example.com",
		Password: "password",
		Admin:    false,
	},
}

var UserResponseTestData = []entity.UserResponse{
	{
		ID:    1,
		Name:  "admin",
		Email: "admin@example.com",
		Admin: true,
	},
	{
		ID:    2,
		Name:  "user01",
		Email: "user01@example.com",
		Admin: false,
	},
}

func GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	switch email {
	case "admin@example.com":
		user = UserTestData[0]
	case "user01@example.com":
		user = UserTestData[1]
	default:
		return nil, errors.New("user not found")
	}
	return &user, nil
}

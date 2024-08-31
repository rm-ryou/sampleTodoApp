package testdata

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

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
		ID:       1,
		Name:     "admin",
		Email:    "admin@example.com",
		Admin:    true,
	},
	{
		ID:       2,
		Name:     "user01",
		Email:    "user01@example.com",
		Admin:    false,
	},
}

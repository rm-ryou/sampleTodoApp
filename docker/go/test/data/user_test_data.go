package testdata

import "github.com/rm-ryou/sampleTodoApp/internal/entity"

var UserTestData = []entity.User{
	{
		ID:       1,
		Name:     "admin",
		Email:    "admin@example.com",
		Password: "$2a$10$evMrGNMf75Yn.C6cplFLie48l3Q6jvvds2ym.ZBJVDlY5rVctA2Qy",
		Admin:    true,
	},
	{
		ID:       2,
		Name:     "user01",
		Email:    "user01@example.com",
		Password: "$2a$10$PYQHWYmxga719jfyqYxWZ.NjdfCo9StyDEdLmemEVjT6sNHGmZlJS",
		Admin:    false,
	},
}

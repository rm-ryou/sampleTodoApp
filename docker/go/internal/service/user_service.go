package service

import (
	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	"github.com/rm-ryou/sampleTodoApp/internal/storage/mysql/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServicer interface {
	GetUser(id int) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	EditUser(id int, password, name, email, Newpassword string) (*entity.User, error)
	DeleteUser(id int) error
}

type UserService struct {
	r repository.IUserRepository
}

func NewUserService(r repository.IUserRepository) *UserService {
	return &UserService{r}
}

func (us *UserService) GetUser(id int) (*entity.User, error) {
	return us.r.ReadUser(id)
}

func (us *UserService) GetUsers() ([]entity.User, error) {
	return us.r.ReadUsers()
}

func setNewPassword(user *entity.User, password, newPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}

	hashed, err := EncryptPassword(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashed

	return nil
}

func (us *UserService) EditUser(id int, password, name, email, newPassword string) (*entity.User, error) {
	user, err := us.r.ReadUser(id)
	if err != nil {
		return nil, err
	}

	if newPassword != "" {
		if err := setNewPassword(user, password, newPassword); err != nil {
			return nil, err
		}
	}

	if user.Name != name {
		user.Name = name
	}
	if user.Email != email {
		user.Email = email
	}

	if err := us.r.UpdateUser(user); err != nil {
		return nil, err
	}

	return us.r.ReadUser(id)
}

func (us *UserService) DeleteUser(id int) error {
	return us.r.DeleteUser(id)
}

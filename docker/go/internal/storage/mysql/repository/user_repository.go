package repository

import (
	"database/sql"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
)

type IUserRepository interface {
	CreateUser(user *entity.User) error
	ReadUser(id int) (*entity.User, error)
	ReadUserByEmail(email string) (*entity.User, error)
	ReadUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id int) error
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *entity.User) error {
	query := `
		INSERT INTO users
			(name, email, password)
		VALUES
			(?, ?, ?)
	`

	stmt, err := ur.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) ReadUser(id int) (*entity.User, error) {
	var user entity.User
	query := `
		SELECT *
		FROM users
		WHERE id = ?
	`

	if err := ur.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Admin,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) ReadUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `
		SELECT *
		FROM users
		WHERE email = ?
	`

	if err := ur.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Admin,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) ReadUsers() ([]entity.User, error) {
	var users []entity.User
	query := `
		SELECT *
		FROM users
	`

	rows, err := ur.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Admin,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(user *entity.User) error {
	query := `
		UPDATE users
		SET
			name = ?,
			email = ?,
			password = ?
		WHERE id = ?
	`

	stmt, err := ur.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Email, user.Password, user.ID); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	query := `
		DELETE FROM users
		WHERE id = ?
	`

	if _, err := ur.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

package auth

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrorUserAlreadyExists = errors.New("User already exists.")
	ErrorUserNotFound      = errors.New("User does not exist.")
	ErrorWrongPassword     = errors.New("Wrong password")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) userExist(ctx context.Context, email string) (bool, error) {
	var exist int
	err := r.db.QueryRowContext(
		ctx,
		"SELECT COUNT(1) FROM users WHERE email = $1",
		email,
	).Scan(&exist)
	if err != nil {
		return exist != 0, err
	}
	return exist != 0, nil
}

func (r *Repository) RegUser(ctx context.Context, user UserProfile) error {
	f, err := r.userExist(ctx, user.Email)
	if err != nil || !f {
		return err
	}
	if !f {
		_, err := r.db.ExecContext(
			ctx,
			"insert into users (name, lastname, email, password) values ($1, $2, $3, $4)",
			user.Name,
			user.Lastname,
			user.Email,
			user.Password,
		)
		if err != nil {
			return err
		}
	}
	return ErrorUserAlreadyExists
}

func (r *Repository) LoginUser(ctx context.Context, user UserProfile) error {
	f, err := r.userExist(ctx, user.Email)
	if err != nil {
		return err
	}

	if !f {
		return ErrorUserNotFound
	}

	var password string
	if err := r.db.QueryRowContext(
		ctx,
		"select password from users where email = $1",
		user.Email,
	).Scan(&password); err != nil {
		return err
	}

	if user.Password == password {
		return nil
	}

	return ErrorWrongPassword
}

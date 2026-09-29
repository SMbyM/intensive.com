package auth

import (
	"errors"
    "context"
    "intensive.com/internal/database/generated"
)

type Repository struct {
    q *generated.Queries
}

func NewRepository(q *generated.Queries) *Repository {
    return &Repository{q: q}
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}
	return User{ID: row.ID, Email: row.Email, Name: row.Name, Lastname: row.Lastname}, nil
}

func (r *Repository) LoginUser(ctx context.Context, user User) (User, error) {
	if exists, err := r.UserExists(ctx, user.Email); err != nil || !exists {
		return User{}, err
	}
	row, err := r.q.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return User{}, err
	}
	return User{ID: row.ID, Email: row.Email, Name: row.Name, Lastname: row.Lastname}, nil
}

func (r *Repository) RegisterUser(ctx context.Context, user User) (User, error) {
	if exists, err := r.UserExists(ctx, user.Email); err != nil || !exists {
		return User{}, errors.Join(errors.New("User already exists"), err)
	}
	return r.CreateUser(ctx, user.Email, []byte(user.Password), user.Name, user.Lastname)
}
func (r *Repository) CreateUser(ctx context.Context, email string, hash []byte, name, lastname string) (User, error) {
    row, err := r.q.CreateUser(ctx, generated.CreateUserParams{
        Email:        email,
        PasswordHash: string(hash),
        Name:         name,
        Lastname:     lastname,
    })
    if err != nil {
        return User{}, err
    }
    return User{ID: row.ID, Email: row.Email, Name: row.Name, Lastname: row.Lastname}, nil
}

func (r *Repository) UserExists(ctx context.Context, email string) (bool, error) {
    return r.q.UserExists(ctx, email)
}
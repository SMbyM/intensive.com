package account

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"intensive.com/internal/database/generated"
)

type Repository struct {
	db *generated.Queries
}

func NewRepository(db *generated.Queries) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(ctx context.Context, uid int32) (User, error) {
	u, err := r.db.GetUserById(ctx, uid)

	if err != nil {
		return User{}, fmt.Errorf("get user by id: %w", err)
	}
	return User{
		ID: u.ID,
		Name: u.Name,
		Lastname: u.Lastname, 
		Nickname: *u.Nickname, 
		Email: u.Email,
		Phone: *u.Phone, 
		Male: *u.Male, 
		Birthday: u.Birthday.Time.String(),
		}, nil
}

func (r *Repository) UpdateUserProfile(ctx context.Context, user User) error {
	time, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		return fmt.Errorf("parse birthday: %w", err)
	}
	birthday := pgtype.Date{Time: time, Valid: true}
	if e := r.db.UpdateUserProfile(ctx, generated.UpdateUserProfileParams{
		ID: user.ID,
		Name: &user.Name,
		Lastname: &user.Lastname,
		Nickname: &user.Nickname,
		Phone: &user.Phone,
		Male: &user.Male,
		Birthday: birthday,
	});
	e != nil {
		return fmt.Errorf("update user profile: %w", err)
	}
	return nil
}

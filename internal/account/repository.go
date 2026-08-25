package account

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(ctx context.Context, uid int) (UserProfile, error) {
	var u UserProfile
	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, lastname, nickname, email, phone, male, birthday FROM users WHERE id = $1",
		uid,
	).Scan(&u.ID, &u.Name, &u.Lastname, &u.Nickname, &u.Email, &u.Phone, &u.Male, &u.Birthday)

	if err != nil {
		return UserProfile{}, fmt.Errorf("get user by id: %w", err)
	}
	return u, nil
}

func (r *Repository) SetFriendship(ctx context.Context, fst int, snd int) error {
	if _, err := r.db.ExecContext(
		ctx,
		"insert into friends values ($1, $2)",
		fst,
		snd,
	); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetFriendList(ctx context.Context, uid int) ([]UserProfile, error) {
	var friends []UserProfile

	rows, err := r.db.QueryContext(ctx, "select snd from friends where fst=$1", uid)
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var snd int
		err := rows.Scan(&snd)
		if err != nil {
			fmt.Println("accountControllers:31")
			return nil, err
		}
		res, err := r.GetUserByID(ctx, snd)
		if err != nil {
			fmt.Println("accountControllers:43")
			return nil, err
		}
		friends = append(friends, res)
	}
	return friends, nil
}

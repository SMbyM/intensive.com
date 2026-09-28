package friends

import (
	"context"

	"intensive.com/internal/database/generated"
)

type Repository struct {
	db *generated.Queries
}

func NewRepository(db *generated.Queries) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateFriendship(ctx context.Context, fst int32, snd int32) error {
	if err := r.db.CreateFriendship(
		ctx,
		generated.CreateFriendshipParams{
			Fst: fst,
			Snd: snd,
		},
	); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFriendship(ctx context.Context, fst int32, snd int32) error {
	if err := r.db.DeleteFriendship(
		ctx,
		generated.DeleteFriendshipParams{
			Fst: fst,
			Snd: snd,
		},
	); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetFriendList(ctx context.Context, uid int32) ([]int32, error) {
	uids, err := r.db.GetFriendList(ctx, uid)
	if err != nil {
		return nil, err
	}
	return uids, nil
}

func (r *Repository) AreFriends(ctx context.Context, fst int32, snd int32) (bool, error) {
	areFriends, err := r.db.AreFriends(ctx, generated.AreFriendsParams{
		Fst: fst,
		Snd: snd,
	})
	if err != nil {
		return false, err
	}
	return areFriends, nil
}
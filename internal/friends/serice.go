package friends

import (
	"context"
	"errors"
)

var ErrInvalidID = errors.New("invalid id")

type friendsRepository interface {
	CreateFriendship(ctx context.Context, fst int32, snd int32) error
	DeleteFriendship(ctx context.Context, fst int32, snd int32) error
	GetFriendList(ctx context.Context, uid int32) ([]int32, error)
	AreFriends(ctx context.Context, fst int32, snd int32) (bool, error)
}

type Service struct {
	repo friendsRepository
}

func NewService(repo friendsRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateFriendship(ctx context.Context, fst int32, snd int32) error {
	if fst <= 0 || snd <= 0 {
		return ErrInvalidID
	}
	return s.repo.CreateFriendship(ctx, fst, snd)
}

func (s *Service) DeleteFriendship(ctx context.Context, fst int32, snd int32) error {
	if fst <= 0 || snd <= 0 {
		return ErrInvalidID
	}
	return s.repo.DeleteFriendship(ctx, fst, snd)
}

func (s *Service) GetFriendList(ctx context.Context, uid int32) ([]int32, error) {
	if uid <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.GetFriendList(ctx, uid)
}

func (s *Service) AreFriends(ctx context.Context, fst int32, snd int32) (bool, error) {
	if fst <= 0 || snd <= 0 {
		return false, ErrInvalidID
	}
	return s.repo.AreFriends(ctx, fst, snd)
}

package account

import (
	"context"
	"errors"
)

var ErrInvalidID = errors.New("invalid user id")

type userRepository interface {
	GetUserByID(ctx context.Context, id int) (UserProfile, error)
	SetFriendship(ctx context.Context, fst int, snd int) error
	GetFriendList(ctx context.Context, uid int) ([]UserProfile, error)
}

type Service struct {
	repo userRepository
}

func NewService(repo userRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProfile(ctx context.Context, uid int) (UserProfile, error) {
	if uid <= 0 {
		return UserProfile{}, ErrInvalidID
	}
	return s.repo.GetUserByID(ctx, uid)
}

func (s *Service) SetFriendship(ctx context.Context, fst int, snd int) error {
	if fst <= 0 || snd <= 0 {
		return ErrInvalidID
	}
	return s.repo.SetFriendship(ctx, fst, snd)
}

func (s *Service) GetFriendList(ctx context.Context, uid int) ([]UserProfile, error) {
	if uid <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.GetFriendList(ctx, uid)
}

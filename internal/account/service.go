package account

import (
	"context"
	"errors"
)

var ErrInvalidID = errors.New("invalid user id")

type userRepository interface {
	GetUserByID(ctx context.Context, id int32) (User, error)
	UpdateUserProfile(ctx context.Context, user User) error
}

type Service struct {
	repo userRepository
}

func NewService(repo userRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProfile(ctx context.Context, uid int32) (User, error) {
	if uid <= 0 {
		return User{}, ErrInvalidID
	}
	return s.repo.GetUserByID(ctx, uid)
}
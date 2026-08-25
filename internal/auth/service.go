package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"regexp"
)

var (
	ErrorInvalidEmail = errors.New("Invalid email.")
)

type authRepository interface {
	RegUser(ctx context.Context, user UserProfile) error
	LoginUser(ctx context.Context, user UserProfile) error
}

type Service struct {
	repo authRepository
}

func NewService(r authRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) RegUser(ctx context.Context, user UserProfile) error {
	if f, err := regexp.MatchString(
		`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
		user.Email,
	); !f || err != nil {
		return errors.Join(ErrorInvalidEmail, err)
	}

	h := sha256.New()
	h.Write([]byte(user.Password))

	hash := hex.EncodeToString(h.Sum(nil))
	user.Password = hash

	return s.repo.RegUser(ctx, user)
}

func (s *Service) LoginUser(ctx context.Context, user UserProfile) error {
	if f, err := regexp.MatchString(
		`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
		user.Email,
	); !f || err != nil {
		return errors.Join(ErrorInvalidEmail, err)
	}

	h := sha256.New()
	h.Write([]byte(user.Password))

	hash := hex.EncodeToString(h.Sum(nil))
	user.Password = hash

	return s.repo.LoginUser(ctx, user)
}

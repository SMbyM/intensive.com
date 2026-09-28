package friends

import (
	"intensive.com/internal/database/generated"
)

func New(db *generated.Queries) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service)
}

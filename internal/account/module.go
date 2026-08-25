package account

import "database/sql"

func New(db *sql.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service)
}

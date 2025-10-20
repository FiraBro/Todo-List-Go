package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Create(context.Context) error
	}
	User interface {
		Create(context.Context) error
	}
}

func NewSQLStorage(db *sql.DB) Storage {
	return Storage{
		Post: &postsSoter{db},
		User: &userStore{db},
	}
}
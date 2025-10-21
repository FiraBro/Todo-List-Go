package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Create(context.Context, *post) error
	}
	User interface {
		Create(context.Context,*users) error
	}
}

func NewSQLStorage(db *sql.DB) Storage {
	return Storage{
		Post: &postsSoter{db},
		User: &userStore{db},
	}
}
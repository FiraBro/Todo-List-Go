package store

import (
	"context"
	"database/sql"
)

type postsSoter struct{
	db *sql.DB
}

func (s *postsSoter) Create(ctx context.Context) error {
	return nil
}
package store

import (
	"context"
	"database/sql"
)
type userStore struct{
	db *sql.DB
}

func (s *userStore) Create(ctx context.Context) error {
	return nil
}
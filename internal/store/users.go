package store

import (
	"context"
	"database/sql"
)
type user struct{
Name string `json:"name"`
Age int8 `json:"age"`
Password string `json:"password"`
CreatedAt string `json:"created_at"`
UpdatedAt string `json:"updated_at"`

}
type userStore struct{
	db *sql.DB
}

func (s *userStore) Create(ctx context.Context,user *user) error {
	query:=`
	INSERT INTO user (user.name,user.password,user.age)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`
	err:=s.db.QueryRowContext(ctx,query,user.Name,user.Age,user.Password).Scan(
		&user.CreatedAt,
		&user.UpdatedAt)
	if err !=nil{
		return err
	}
	return nil
}
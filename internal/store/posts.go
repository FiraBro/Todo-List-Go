package store

import (
	"context"
	"database/sql"
)

type post struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserID int64 `json:"user_id"`
	// Tags []string `json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type postsSoter struct{
	db *sql.DB
}


func (s *postsSoter) Create(ctx context.Context,post *post) error {
	query:=`
	INSERT INTO post (title, content, user_id)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`
	err:= s.db.QueryRowContext(ctx, query, post.Title, post.Content, post.UserID).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil{
		return err
	}
	return nil
}
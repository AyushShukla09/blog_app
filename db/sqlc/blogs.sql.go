// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: blogs.sql

package db

import (
	"context"
	"time"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blog_posts(title,description,body)
VALUES($1,$2,$3)
RETURNING id, title, description, body, created_at, updated_at
`

type CreateBlogParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (BlogPost, error) {
	row := q.db.QueryRowContext(ctx, createBlog, arg.Title, arg.Description, arg.Body)
	var i BlogPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blog_posts
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getAllBlog = `-- name: GetAllBlog :many
SELECT id, title, description, body, created_at, updated_at FROM blog_posts
`

func (q *Queries) GetAllBlog(ctx context.Context) ([]BlogPost, error) {
	rows, err := q.db.QueryContext(ctx, getAllBlog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BlogPost
	for rows.Next() {
		var i BlogPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBlog = `-- name: GetBlog :one
SELECT id, title, description, body, created_at, updated_at FROM blog_posts WHERE id = $1
`

func (q *Queries) GetBlog(ctx context.Context, id int64) (BlogPost, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i BlogPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateBlog = `-- name: UpdateBlog :one
UPDATE blog_posts
SET
    title = $1,
    description = $2,
    body = $3,
    updated_at = $4
WHERE
    id = $5
RETURNING id, title, description, body, created_at, updated_at
`

type UpdateBlogParams struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	UpdatedAt   time.Time `json:"updated_at"`
	ID          int64     `json:"id"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (BlogPost, error) {
	row := q.db.QueryRowContext(ctx, updateBlog,
		arg.Title,
		arg.Description,
		arg.Body,
		arg.UpdatedAt,
		arg.ID,
	)
	var i BlogPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

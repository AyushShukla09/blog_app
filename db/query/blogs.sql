-- name: CreateBlog :one
INSERT INTO blog_posts(title,description,body)
VALUES($1,$2,$3)
RETURNING *;

-- name: GetAllBlog :many
SELECT * FROM blog_posts;

-- name: GetBlog :one
SELECT * FROM blog_posts WHERE id = $1;

-- name: DeleteBlog :exec
DELETE FROM blog_posts
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blog_posts
SET
    title = $1,
    description = $2,
    body = $3,
    updated_at = $4
WHERE
    id = $5
RETURNING *;

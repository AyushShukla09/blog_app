package db

import "context"

type Repo interface {
	CreateBlog(ctx context.Context, arg CreateBlogParams) (BlogPost, error)
	GetAllBlog(ctx context.Context) ([]BlogPost, error)
	GetBlog(ctx context.Context, id int64) (BlogPost, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) (BlogPost, error)
	DeleteBlog(ctx context.Context, id int64) error
}

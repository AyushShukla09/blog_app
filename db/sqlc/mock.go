package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateBlog(ctx context.Context, arg CreateBlogParams) (BlogPost, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(BlogPost), args.Error(1)
}
func (m *MockRepo) GetAllBlog(ctx context.Context) ([]BlogPost, error) {
	args := m.Called(ctx)
	return args.Get(0).([]BlogPost), args.Error(1)
}
func (m *MockRepo) GetBlog(ctx context.Context, id int64) (BlogPost, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(BlogPost), args.Error(1)
}
func (m *MockRepo) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (BlogPost, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(BlogPost), args.Error(1)
}
func (m *MockRepo) DeleteBlog(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

package db

import (
	"blog_app/utils"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomBlog(t *testing.T) BlogPost {
	args := CreateBlogParams{
		Title:       utils.RandomStringGenerator(10),
		Description: utils.RandomStringGenerator(20),
		Body:        utils.RandomStringGenerator(15),
	}

	blog, err := testQueries.CreateBlog(context.Background(), args)
	require.NoError(t, err)
	require.NotZero(t, blog.ID)
	require.Equal(t, args.Title, blog.Title)
	require.Equal(t, args.Description, blog.Description)
	require.Equal(t, args.Body, blog.Body)
	return blog
}

func TestCreateBlog(t *testing.T) {
	createRandomBlog(t)
}

func TestGetAllBlog(t *testing.T) {
	createRandomBlog(t)
	blogs, err := testQueries.GetAllBlog(context.Background())
	require.NoError(t, err)
	require.NotZero(t, len(blogs))
	require.NotEmpty(t, blogs)
}

func TestGetBlog(t *testing.T) {
	blog1 := createRandomBlog(t)
	blog2, err := testQueries.GetBlog(context.Background(), blog1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, blog2)
	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.Description, blog2.Description)
	require.Equal(t, blog1.Body, blog2.Body)

}

func TestDeleteBlog(t *testing.T) {
	blog1 := createRandomBlog(t)
	err := testQueries.DeleteBlog(context.Background(), blog1.ID)
	require.NoError(t, err)
	blog2, err := testQueries.GetBlog(context.Background(), blog1.ID)
	require.Error(t, err)
	require.Empty(t, blog2)
	require.Zero(t, blog2.ID)
	require.NotEqual(t, blog1.ID, blog2.ID)
	require.NotEqual(t, blog1.Title, blog2.Title)
	require.NotEqual(t, blog1.Description, blog2.Description)
	require.NotEqual(t, blog1.Body, blog2.Body)
}

func TestUpdateBlog(t *testing.T) {
	blog1 := createRandomBlog(t)
	args := UpdateBlogParams{
		Title:       utils.RandomStringGenerator(10),
		Description: utils.RandomStringGenerator(20),
		Body:        utils.RandomStringGenerator(15),
		UpdatedAt:   time.Now(),
		ID:          blog1.ID,
	}
	newBlog, err := testQueries.UpdateBlog(context.Background(), args)
	require.NoError(t, err)
	require.NotZero(t, newBlog.ID)
	require.Equal(t, blog1.ID, newBlog.ID)
	require.NotEqual(t, blog1.Title, newBlog.Title)
	require.NotEqual(t, blog1.Description, newBlog.Description)
	require.NotEqual(t, blog1.Body, newBlog.Body)

}

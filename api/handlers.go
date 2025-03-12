package api

import (
	db "blog_app/db/sqlc"
	response "blog_app/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Blog struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateBlog struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (s *Server) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	newBlog, err := s.queries.CreateBlog(context.Background(), db.CreateBlogParams{
		Title:       blog.Title,
		Description: blog.Description,
		Body:        blog.Body,
	})
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Send(w, http.StatusCreated, newBlog)
}

func (s *Server) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := s.queries.GetAllBlog(context.Background())
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Send(w, http.StatusOK, blogs)
}

func (s *Server) GetBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blogID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	blog, err := s.queries.GetBlog(context.Background(), blogID)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Send(w, http.StatusOK, blog)
}

func (s *Server) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blogID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	if err = s.queries.DeleteBlog(context.Background(), blogID); err != nil {
		response.Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Send(w, http.StatusOK, "Blog successfully deleted")
}

type UpdateBlog struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (s *Server) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blogID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	var blog UpdateBlog
	if err = json.NewDecoder(r.Body).Decode(&blog); err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	updatedBlog, err := s.queries.UpdateBlog(context.Background(), db.UpdateBlogParams{
		Title:       blog.Body,
		Description: blog.Description,
		Body:        blog.Body,
		ID:          blogID,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Send(w, http.StatusOK, updatedBlog)
}

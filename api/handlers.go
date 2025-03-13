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

// @Summary Add blog
// @Description Endpoint to create a new blog
// @Tags Blog
// @Accept json
// @Produce json
// @Param post body CreateBlog true "Request Body to create a blog post"
// @Success 201 {object} Blog "Successful Response"
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 500 {object} response.Response "Internal Server Error"
// @Router /blog [post]
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

// @Summary List blogs
// @Description Endpoint to fetch all blogs
// @Tags Blogs
// @Accept json
// @Produce json
// @Success 200 {array} Blog "Successful Response"
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 500 {object} response.Response "Internal Server Error"
// @Router /blogs [get]
func (s *Server) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := s.queries.GetAllBlog(context.Background())
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Send(w, http.StatusOK, blogs)
}

// @Summary Show blog
// @Description Endpoint to fetch blog by ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} Blog "Successful Response"
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 500 {object} response.Response "Internal Server Error"
// @Router /blog/{id} [get]
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

// @Summary Delete a blog
// @Description Endpoint to delete blog by ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} string "Successful Response"
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 500 {object} response.Response "Internal Server Error"
// @Router /blog/{id} [delete]
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

// @Summary Update a blog
// @Description Endpoint to update blog by ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} Blog "Successful Response"
// @Failure 400 {object} response.Response "Bad Request"
// @Failure 500 {object} response.Response "Internal Server Error"
// @Router /blog/{id} [put]
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

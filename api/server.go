package api

import (
	db "blog_app/db/sqlc"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	queries *db.Queries
	router  *chi.Mux
}

func NewServer(db *db.Queries) *chi.Mux {
	router := chi.NewRouter()
	server := &Server{
		queries: db,
		router:  router,
	}
	router.Post("/blog", server.CreateBlog)
	router.Get("/blogs", server.GetAllBlogs)
	router.Get("/blog/{id}", server.GetBlog)
	router.Delete("/blog/{id}", server.DeleteBlog)
	router.Put("/blog/{id}", server.UpdateBlog)
	return server.router
}

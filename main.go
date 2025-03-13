package main

import (
	db "blog_app/db/sqlc"
	"database/sql"
	"log"
	"net/http"

	"blog_app/api"

	_ "blog_app/docs"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const (
	dbdriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/blog_post?sslmode=disable"
)

// @title Blog Post API
// @version 1.0
// @description This is a sample blog CRUD API server.
// @host localhost:8080
// @BasePath /
// @contact.name Ayush Shukla
// @contact.email ayush.shukla8797@gmail.com
func main() {
	conn, err := sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	// create db connection object
	queries := db.New(conn)
	// create router
	router := api.NewServer(queries)
	router.Get("/swagger/*", httpSwagger.WrapHandler)
	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

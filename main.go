package main

import (
	db "blog_app/db/sqlc"
	"database/sql"
	"log"
	"net/http"

	"blog_app/api"

	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/blog_post?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	// create db connection object
	queries := db.New(conn)
	// create router
	router := api.NewServer(queries)

	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

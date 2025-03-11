package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbdriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/blog_post?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}

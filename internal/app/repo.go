package app

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// DB hold the Database instace.
type DB struct {
	*sql.DB
}

// Factory method which create a new DB.
func NewDB() *DB {
	return &DB{
		connectDatabase("postgres://user:pass@localhost:5432/sample?sslmode=disable"),
	}
}

// connectDatabase is a DB connection helper function.
func connectDatabase(connectionURL string) *sql.DB {
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

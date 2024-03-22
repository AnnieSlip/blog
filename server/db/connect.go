package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5111
	user   = "anniejavakhishvili"
	dbname = "postgres"
)

func NewDB() (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the PostgreSQL database!")
	return db, nil
}

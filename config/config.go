package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func InitDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, error := sql.Open("postgres", dsn)

	if error != nil {
		log.Fatal("DB connection error:", error)
	}

	if error = db.Ping(); error != nil {
		log.Fatal("DB ping error:", error)
	}

	return db
}

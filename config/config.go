package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5433         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
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

package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	// Optional: verify the connection
	if err := db.Ping(); err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	DB = db
	fmt.Println("Successfully connected to the database")
}

package database

import (
	"database/sql"
	"log"
	"os"
)

func GetDBConnection() *sql.DB {
	// Get the database connection
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

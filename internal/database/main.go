package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	log.Print("Initializing SQLITE3 database...")

	db, err := sql.Open("sqlite3", "../../db/project.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

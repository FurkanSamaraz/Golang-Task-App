package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func Openconnention() *sql.DB {

	db, err := sql.Open("sqlite3", "Database.db")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

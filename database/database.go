package database

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, errDb := sql.Open("postgres", "host=db port=5432 user=root password=root dbname=rest sslmode=disable")
	if errDb != nil {
		log.Fatal(errDb)
	}
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

package pg

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, errDB := sql.Open("postgres", "host=db port=5432 user=root password=root dbname=rest sslmode=disable")
	if errDB != nil {
		log.Fatal(errDB)
	}

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

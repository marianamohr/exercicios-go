package db

import (
	"database/sql"
	"log"
)

func ConectaDb() *sql.DB {
	connStr := "user=postgres dbname=alura_loja password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

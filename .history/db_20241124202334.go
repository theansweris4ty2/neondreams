package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func dbOpen() {
	connStr := "postgres://postgres:babbage@localhost:8080/neondreams?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}

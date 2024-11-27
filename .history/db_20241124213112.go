package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func dbOpen() *sql.DB {
	connStr := "postgres://user:babbage@localhost:8080/neondreams?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db

}
func createMovieTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS movie (
	title STRING VARCHAR(100) NOT NULL,
	director STRING VARCHAR(100) NOT NULL,
	genre STRING
	available BOOLEAN,
	created timestamp DEFAULT NOW())`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

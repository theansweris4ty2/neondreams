package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)
const (
	host = "localhost"
	port = 8080
	user = "postgres"
	password = "babbage"
	dbname = "neondreams"
)

func dbOpen() *sql.DB {
	psqlInfo := fmt.Sprintf"postgres://user:babbage@localhost:8080/neondreams?sslmode=disable"
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

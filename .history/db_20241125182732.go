package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5342
	user     = "postgres"
	password = "babbage"
	dbname   = "neondreams"
)

func dbOpen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

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

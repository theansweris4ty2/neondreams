package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Movie struct {
	name     string
	director string
}

func openDb() *sql.DB {

	connStr := "postgresql://theansweris4ty2:5f2VrRrbbPozGkDMkruV6NNcUfx6fTn8@dpg-ct37s5i3esus73f31t3g-a.oregon-postgres.render.com/neondreamsdb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("there was an error %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("There was an error: %v", err)
	}
	fmt.Println("Successfully connected to db")
	return db
}

func createMovieTable(db *sql.DB) {
	query := `CREATE FOREIGN TABLE IF NOT EXISTS movie(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	director VARCHAR(100) NOT NULL, 
	created timestamp DEFAULT NOW()
	)`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("error creating table")
		log.Fatal(err)
	}
}
func insertMovie(db *sql.DB, movie Movie) int {
	query := `INSERT INTO movie (name, director)
	VALUES($1, $2) RETURNING id`

	var pk int
	err := db.QueryRow(query, movie.name, movie.director).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
func getMovie(db *sql.DB, pk int) {
	var name string
	var director string
	query := `SELECT name, director FROM movie WHERE id = $1`
	err := db.QueryRow(query, pk).Scan(&name, &director)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The movie you selected is %s directed by %s", name, director)
}

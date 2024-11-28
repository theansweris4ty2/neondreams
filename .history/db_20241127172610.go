package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Movie struct {
	Name     string
	Director string
}
type Book struct {
	Title  string
	Author string
}

func openDb() *sql.DB {

	connStr := "postgresql://theansweris4ty2:5f2VrRrbbPozGkDMkruV6NNcUfx6fTn8@dpg-ct37s5i3esus73f31t3g-a.oregon-postgres.render.com/neondreamsdb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("there was an error %v", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("There was an error: %v", err)
	}
	fmt.Println("Successfully connected to db")
	return db
}

func createBookTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS book(
	id SERIAL PRIMARY KEY,
	Title VARCHAR(100) NOT NULL,
	Author VARCHAR(100) NOT NULL, 
	created timestamp DEFAULT NOW()
	)`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("error creating table")
		log.Fatal(err)
	}
}
func insertBook(db *sql.DB, book Book) int {
	query := `INSERT INTO book (title, author)
	VALUES($1, $2) RETURNING id`

	var pk int
	err := db.QueryRow(query, book.Title, book.Author).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
func getAllBooks(db *sql.DB) ([]Book, error) {
	var books []Book
	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Title, &book.Author)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return books, err
	}
	for _, b := range books {
		fmt.Printf("The title of this book is %s", b.Title)
	}
	return books, nil
}

func getBook(db *sql.DB, pk int) (string, string) {
	var title string
	var author string
	query := `SELECT name, director FROM movie WHERE id = $1`
	err := db.QueryRow(query, pk).Scan(&title, &author)
	if err != nil {
		log.Fatal(err)
	}
	return title, author
}

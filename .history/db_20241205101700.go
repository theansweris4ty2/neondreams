package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// TODO Add genre field to all structs and a review field to Movie, Book, and TV

type Movie struct {
	Title    string
	Director string
}
type Book struct {
	Title  string
	Author string
}
type Show struct {
	Title string
	Genre string
}
type Article struct {
	Title  string
	Author string
	Blog   string
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
func createMovieTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS movie(
	id SERIAL PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Director VARCHAR(100) NOT NULL, 
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
func insertMovie(db *sql.DB, movie Movie) int {
	query := `INSERT INTO movie (name, director)
	VALUES($1, $2) RETURNING id`

	var pk int
	err := db.QueryRow(query, movie.Title, movie.Director).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
func getAllBooks(db *sql.DB) ([]Book, error) {
	var books []Book
	rows, err := db.Query("SELECT title, author FROM book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var author string
		err := rows.Scan(&title, &author)
		if err != nil {
			return books, err
		}
		book := Book{title, author}
		books = append(books, book)

		if err = rows.Err(); err != nil {
			return books, err
		}
	}

	return books, nil
}
func getAllMovies(db *sql.DB) ([]Movie, error) {
	var movies []Movie
	rows, err := db.Query("SELECT title, director FROM movies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var director string
		err := rows.Scan(&title, &director)
		if err != nil {
			return movies, err
		}
		movie := Movie{title, director}
		movies = append(movies, movie)

		if err = rows.Err(); err != nil {
			return movies, err
		}
	}

	return movies, nil
}
func getAllArticles(db *sql.DB) ([]Article, error) {
	var articles []Article
	rows, err := db.Query("SELECT title, author FROM articles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var author string
		var blog string
		err := rows.Scan(&title, &author, &blog)
		if err != nil {
			return articles, err
		}
		article := Article{title, author, blog}
		articles = append(articles, article)

		if err = rows.Err(); err != nil {
			return articles, err
		}
	}

	return articles, nil
}
func getAllShows(db *sql.DB) ([]Show, error) {
	var shows []Show
	rows, err := db.Query("SELECT title, genre FROM shows")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var genre string
		err := rows.Scan(&title, &genre)
		if err != nil {
			return shows, err
		}
		show := Show{title, genre}
		shows = append(shows, show)

		if err = rows.Err(); err != nil {
			return shows, err
		}
	}

	return shows, nil
}

func getBook(db *sql.DB, t string) (string, string) {
	var title string
	var author string
	query := `SELECT title, author FROM book WHERE title = $1`
	err := db.QueryRow(query, t).Scan(&title, &author)
	if err != nil {
		log.Fatal(err)
	}
	return title, author
}
func getMovie(db *sql.DB, t string) (string, string) {
	var title string
	var director string
	query := `SELECT name, director FROM movie WHERE name = $1`
	err := db.QueryRow(query, t).Scan(&title, &director)
	if err != nil {
		log.Fatal(err)
	}
	return title, director
}
func getArticle(db *sql.DB, t string) (string, string, string) {
	var title string
	var author string
	var blog string
	query := `SELECT title, author, blog FROM articles WHERE title = $1`
	err := db.QueryRow(query, t).Scan(&title, &author, &blog)
	if err != nil {
		log.Fatal(err)
	}
	return title, author, blog
}
func getShow(db *sql.DB, t string) (string, string) {
	var title string
	var genre string
	query := `SELECT title, genre FROM show WHERE title = $1`
	err := db.QueryRow(query, t).Scan(&title)
	if err != nil {
		log.Fatal(err)
	}
	return title, genre
}

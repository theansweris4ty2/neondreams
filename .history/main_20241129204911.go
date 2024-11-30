package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-book/", bookformHandler)
	http.HandleFunc("/add-movie/", movieformHandler)
	http.HandleFunc("/books/", booksHandler)
	http.HandleFunc("/movies/", moviesHandler)
	http.HandleFunc("/articles/", articlesHandler)
	http.HandleFunc("/book-list/", booklistHandler)
	http.HandleFunc("/movie-list/", movielistHandler)
	http.HandleFunc("/article-list/", articlelistHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "articles.html", nil)
}
func bookformHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	htmlStr := fmt.Sprintf("<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s - %s</li>", title, author)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
	db := openDb()
	defer db.Close()
	var newBook = Book{title, author}
	createBookTable(db)
	pk := insertBook(db, newBook)
	getBook(db, pk)
}
func movieformHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	htmlStr := fmt.Sprintf("<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s - %s</li>", title, director)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
	db := openDb()
	defer db.Close()
	var newMovie = Movie{title, director}
	createMovieTable(db)
	pk := insertMovie(db, newMovie)
	getMovie(db, pk)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "books.html", nil)
}
func moviesHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "movies.html", nil)
}
func booklistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	books, _ := getAllBooks(db)

	tpl, _ := template.New("t").Parse(`
		<ul>
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}} - {{.Author}}</li>
		{{end}}
	    </ul>,`)
	tpl.Execute(w, books)

}
func movielistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	movies, _ := getAllMovies(db)

	tpl, _ := template.New("t").Parse(`
		<ul>
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}} - {{.Director}}</li>
		{{end}}
	    </ul>,`)
	tpl.Execute(w, movies)

}
func articlelistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	articles, _ := getAllArticles(db)

	tpl, _ := template.New("t").Parse(`
		<ul>
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}} - {{.Author}}<p>{{.Blog}}</p></li>
		{{end}}
	    </ul>,`)
	tpl.Execute(w, articles)

}
func bookHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	title, author := getBook(db, 1)
	htmlStr := fmt.Sprintf("<ul><li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s - %s</li></ul>", title, author)
	tpl, _ := template.New("t").Parse(htmlStr)

	tpl.Execute(w, nil)
}

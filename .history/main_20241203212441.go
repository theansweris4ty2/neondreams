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
	http.HandleFunc("/books/", booksHandler)
	http.HandleFunc("/movies/", moviesHandler)
	http.HandleFunc("/articles/", articlesHandler)
	http.HandleFunc("/tv/", tvHandler)
	http.HandleFunc("/book-list/", booklistHandler)
	http.HandleFunc("/movie-list/", movielistHandler)
	http.HandleFunc("/articles-list/", articlelistHandler)
	http.HandleFunc("/show-list", showlistHandler)
	http.HandleFunc("/select-book/", selectBookHandler)
	http.HandleFunc("/select-movie/", selectMovieHandler)
	http.HandleFunc("/select-article/", selectArticleHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "articles.html", nil)
}
func booksHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "books.html", nil)
}
func moviesHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "movies.html", nil)
}
func tvHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "tv.html", nil)
}
func selectBookHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	// author := r.PostFormValue("author")
	db := openDb()
	defer db.Close()
	t, a := getBook(db, title)
	htmlStr := fmt.Sprintf("<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s - %s</li>", t, a)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
}
func selectMovieHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	// director := r.PostFormValue("director")
	db := openDb()
	defer db.Close()
	t, d := getMovie(db, title)
	htmlStr := fmt.Sprintf("<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s - %s</li>", t, d)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
}
func selectArticleHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	// author := r.PostFormValue("author")
	db := openDb()
	defer db.Close()
	_, _, b := getArticle(db, title)
	// TODO change the getArticle field to id from title
	htmlStr := fmt.Sprintf(`<div class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> <p>%s</p>
	</div>`, b)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
}
func selectShow(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	db := openDb()
	defer db.Close()
	t, g := getShow(db, title)
	htmlStr := fmt.Sprintf("<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s-%s</li>", t, g)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
}

func booklistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	books, _ := getAllBooks(db)

	tpl, _ := template.New("t").Parse(`
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}} - {{.Author}}</li>
		{{end}}`)
	// TODO add id value
	tpl.Execute(w, books)

}
func movielistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	movies, _ := getAllMovies(db)

	tpl, _ := template.New("t").Parse(`
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}} - {{.Director}}</li>
		{{end}}`)
	tpl.Execute(w, movies)

}
func articlelistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	articles, _ := getAllArticles(db)
	// TODO Change this back to the same form format as the other pages, but have input field for id rather than title or author
	tpl, _ := template.New("t").Parse(`
		{{range .}}
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-px max-w-screen-sm'>{{.Title}} - {{.Author}}<form hx-post="/article/" hx-target="#article-list" hx-swap="innerHTML"><input type="hidden" name="title" value={{.Title}}> <button type="submit"
	class="inline-block w-full rounded bg-gray-700 m-1 px-6 pb-2 pt-2.5 text-xs font-medium uppercase leading-normal text-white shadow-primary-3 transition duration-150 ease-in-out hover:bg-primary-accent-300 hover:shadow-primary-2 focus:bg-primary-accent-300 focus:shadow-primary-2 focus:outline-none focus:ring-0 active:bg-primary-600 active:shadow-primary-2 dark:shadow-black/30 dark:hover:bg-gray-500 shadow-dark-strong dark:focus:shadow-dark-strong dark:active:shadow-dark-strong">Select</button></form></li>
		{{end}}`)

	tpl.Execute(w, articles)
}
func showlistHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	shows, _ := getAllShows(db)

	tpl, _ := template.New("t").Parse(`
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}}{{.Genre}}</li>
		{{end}}`)
	tpl.Execute(w, shows)

}

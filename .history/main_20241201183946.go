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
	http.HandleFunc("/book-list/", booklistHandler)
	http.HandleFunc("/movie-list/", movielistHandler)
	http.HandleFunc("/articles-list/", articlelistHandler)
	http.HandleFunc("/add-book/", selectBookHandler)
	http.HandleFunc("/add-movie/", selectMovieHandler)
	http.HandleFunc("/article/", selectArticleHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "articles.html", nil)
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

	htmlStr := fmt.Sprintf("<div class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> <p>%s</p></div>", b)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)
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
		{{range .}}	
	<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'>{{.Title}} - {{.Author}}</li>
		{{end}}`)
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

	// tpl, _ := template.New("t").Parse(`
	// 	{{range .}}
	// <li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-px max-w-screen-sm'>{{.Title}} - {{.Author}}</li>
	// 	{{end}}`)
	tpl, _ := template.New("t").Parse(`
	{{range .}}
	<form hx-post="/article/" hx-target="#article-box" hx-swap="innerHTML" hx-on::after-request="this.reset()"><div class="grid grid-cols-2 gap-4"><input type="text" placeholder={{.Title}} value="{{.Title}}"class="peer block min-h-[auto] w-full rounded border-0 bg-pink-700 px-3 py-[0.32rem] leading-[1.6] outline-none transition-all duration-200 ease-linear focus:placeholder:opacity-100 peer-focus:text-primary data-[twe-input-state-active]:placeholder:opacity-100 motion-reduce:transition-none dark:text-black dark:placeholder:text-white dark:autofill:shadow-autofill dark:peer-focus:text-primary [&:not([data-twe-input-placeholder-active])]:placeholder:opacity-1">
	<button type="submit" class="inline-block w-full rounded bg-pink-700 px-0 pb-0 m-1 pt-0 text-xs font-medium uppercase leading-normal text-white shadow-primary-3 transition duration-150 ease-in-out hover:bg-primary-accent-300 hover:shadow-primary-2 focus:bg-primary-accent-300 focus:shadow-primary-2 focus:outline-none focus:ring-0 active:bg-primary-600 active:shadow-primary-2 dark:shadow-black/30 dark:hover:bg-pink-500 shadow-dark-strong dark:focus:shadow-dark-strong dark:active:shadow-dark-strong">Select</button>
	</div>
	</form>
	{{end}}`)
	tpl.Execute(w, articles)

}

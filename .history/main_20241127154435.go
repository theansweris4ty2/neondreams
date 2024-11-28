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
	http.HandleFunc("/add-book/", formHandler)
	http.HandleFunc("/film-list/", filmsHandler)
	http.HandleFunc("/list/", http.HandlerFunc(listHandler))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.html", nil)
}
func formHandler(w http.ResponseWriter, r *http.Request) {
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

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "films.html", nil)
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	var books []Book
	books, _ = getAllBooks(db)
	tpl, _ := template.New("t").Parse(`
		<ul>
		{{range .}}
		<li>{{.Title}} - {{.Author}}</li>
		{{end}}
        </ul>,`)
	tpl.Execute(w, books)

}

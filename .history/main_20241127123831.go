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
	http.HandleFunc("/add-film/", formHandler)
	http.HandleFunc("/films-list/", filmsHandler)
	// http.HandleFunc("/database/", http.HandlerFunc(databaseHandler))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.html", nil)
}
func formHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("Id = %d", pk)
	getMovie(db, pk)
}

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	db := openDb()
	defer db.Close()
	var name string
	var director string
	name, director = getMovie(db, 1)
	htmlStr := fmt.Sprintf("<li class='inline-block w-full rounded bg-pink-700 px-6 pb-2.5 pt-2.5 text-lg font-large uppercase leading-normal text-white m-2'> %s - %s</li>", name, director)
	tpl, _ := template.New("t").Parse(htmlStr)
	tpl.Execute(w, nil)

}

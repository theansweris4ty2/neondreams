package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	connStr := "postgres://user:babbage@localhost:8080/neondreams?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-film/", formHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":5342", nil)

	createMovieTable(db)

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
}

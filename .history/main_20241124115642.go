package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type Film struct {
	Title    string
	Director string
}

var films = map[string][]Film{
	"Films": {
		{Title: "Lord of the Rings", Director: "Peter Jackson"},
		{Title: "Big Trouble in Little China", Director: "John Carpenter"},
		{Title: "Star Wars", Director: "George Lucas"},
	},
}

func main() {

	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-film/", formHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.html", films)
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	htmlStr := fmt.Sprintf("<p class='text-sm font-medium text-gray-900 truncate dark:text-black'>%s - %s </p>", title, director)
}

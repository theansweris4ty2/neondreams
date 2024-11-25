package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/form", formHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.html", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout2.html", nil)
}

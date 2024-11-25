package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/navbar", navbarHandler)

	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.hml", nil)
	// tpl.Execute(w, nil)
}

func navbarHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "navbar.html", nil)
}

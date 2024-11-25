package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	http.Handle("/assets", http.StripPrefix("/assets/", http.FileServer(http.Dir("/assets/"))))
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/footer", footerHandler)
	http.HandleFunc("/navbar", navbarHandler)

	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.html", nil)
}

func footerHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout.html", nil)
}
func navbarHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "layout2.html", nil)
}

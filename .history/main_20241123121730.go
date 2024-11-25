package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl, _ = template.ParseGlob("templates/*.go.html")
	// tpl, _ = template.ParseFiles("home.html")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/footer", footerHandler)

	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func footerHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "footer.go.html", nil)
}

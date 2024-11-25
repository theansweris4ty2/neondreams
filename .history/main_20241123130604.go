package main

import (
	"html/template"
	"net/http"
)

// var tpl = template.Must(template.ParseGlob("templates/*"))
var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	// tpl, _ = template.ParseFiles("templates/header.html")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/footer", footerHandler)

	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "header.html", nil)
	// tpl.Execute(w, nil)
}

func footerHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "footer.html", nil)
}

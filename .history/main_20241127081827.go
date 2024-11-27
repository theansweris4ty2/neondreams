package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var tpl *template.Template
var db *sql.DB

func main() {
	connStr := "postgresql://theansweris4ty2:5f2VrRrbbPozGkDMkruV6NNcUfx6fTn8@dpg-ct37s5i3esus73f31t3g-a.oregon-postgres.render.com/neondreamsdb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("there was an error %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("There was an error: %v", err)
	}
	fmt.Println("Successfully connected to db")
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-film/", formHandler)
	http.HandleFunc("/database/", http.HandlerFunc(databaseHandler))
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
}
func databaseHandler(w http.ResponseWriter, r *http.Request) {
	var secondMovie = Movie{"Star Wars", "George Lucas"}
	createMovieTable(db)
	pk := insertMovie(db, secondMovie)
	fmt.Printf("Id = %d", pk)
	getMovie(db, pk)
}

package main

import (
	"net/http"
	"html/template"
)
func main(){
	tpl, _ := http.ParseFile("/template.go.html")
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe("8080", nil)
	func homeHandler(w http.ResponseWriter, r * http.Request){
	
	}
}
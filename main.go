package main

import (
	"html/template"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.ListenAndServe(":3000", routes())
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, nil)
}

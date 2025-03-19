package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	log.Fatal(http.ListenAndServe(port(), routes()))
}

func port() string {
	host := "0.0.0.0"

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return host + ":" + port
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, nil)
}

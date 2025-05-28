package main

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"github.com/gorilla/pat"
	handler "gonote/handler"
	"log"
	"net/http"
)

func main() {

	log.Println("Initializing....")
	static := http.Dir("../..")
	http.Handle("/static/", http.FileServer(static))
	r := pat.New()
	r.Get("/", handler.IndexHandler)
	r.Post("/create", handler.CreateHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":4000", nil))
}

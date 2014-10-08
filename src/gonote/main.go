package main

import (
	"github.com/gorilla/pat"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query()
	name := p.Get(":name")
	w.Write([]byte("hello" + name))
}

func main() {

	r := pat.New()
	r.Get("/", IndexHandler)

	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/", r)

	log.Println("Listen on", ":4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

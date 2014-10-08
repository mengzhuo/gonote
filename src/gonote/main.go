package main

import (
	"db"
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
	http.Handle("/", r)
	log.Println("Listen on ", ":4000")
	http.ListenAndServe(":4000", r)
}

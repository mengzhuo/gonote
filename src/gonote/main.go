package main

import (
	"github.com/gorilla/pat"
	db "gonote/database"
	handler "gonote/handler"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.C("gonote").Count()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(count)
	w.Write([]byte("hello"))
}

func main() {

	log.Println("Initializing....")
	static := http.Dir("../..")
	http.Handle("/static/", http.FileServer(static))
	r := pat.New()
	r.Get("/", IndexHandler)
	r.Post("/create", handler.CreateHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":4000", nil))
}

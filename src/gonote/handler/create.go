package handler

import (
	//db "gonote/database"
	utils "gonote/utils"
	"log"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	log.Println(r.PostForm["Title"])
	utils.WriteJson(w, r.PostForm)
}

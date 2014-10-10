package handler

import (
	db "gonote/database"
	"gonote/utils"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	title := r.PostForm.Get("Title")
	content := r.PostForm.Get("Content")
	tags := r.PostForm["Tags"]

	log.Println("Data:", title, content, tags)

	id := bson.NewObjectId()

	info, err := db.C("gonote").UpsertId(bson.M{"_id": id},
		db.Note{
			Id:        id,
			Title:     title,
			Content:   content,
			Tags:      tags,
			Create_at: time.Now(),
			Update_at: time.Now(),
			Status:    "publish"})

	if err != nil {
		log.Fatal("Create Error:", err)
	}

	log.Println("Result:", info, err)
	utils.WriteJson(w, info)
}

package handler

import (
	db "gonote/database"
	//"gonote/utils"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	//"time"
	"text/template"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	//utils.WriteJson(w, info)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	var results []db.Note

	err := db.C("gonote").Find(bson.M{"status": "publish"}).Sort("-create_at").All(&results)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(results)

	t := template.New("Index")

	tpl, err := t.ParseFiles("index.html")
	tpl.Execute(w, results)
}

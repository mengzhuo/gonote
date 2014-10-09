package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Note struct {
	Id        bson.ObjectId `bson:"_id"`
	Title     string
	Content   string
	Create_at time.Time
	Update_at time.Time
	Status    string
	Tags      []string
}

var session *mgo.Session

func Session() *mgo.Session {

	if session == nil {
		log.Println("Session Init")
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatal("Database is down:%v", err)
			panic("Database OFFLINE")
		}
	}

	return session.Clone()
}

func C(name string) *mgo.Collection {
	session := Session()
	return session.DB("gonote").C(name)
}

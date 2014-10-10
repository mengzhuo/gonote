package utils

import (
	"log"
	"net/http"
)

func RaiseError(w http.ResponseWriter, obj interface{}) {

	log.Fatal(obj)
	d := map[string]interface{}{"err": obj}
	WriteJson(w, d)
}

package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func WriteJson(w http.ResponseWriter, obj interface{}) {

	cov, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(cov)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(cov)
}

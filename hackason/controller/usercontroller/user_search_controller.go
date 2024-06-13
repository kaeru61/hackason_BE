package usercontroller

import (
	"db/application/userapplication"
	"encoding/json"
	"log"
	"net/http"
)

var name string

func UserSearchController(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name = queryParams.Get("name")
	users, err := userapplication.UserSearchApplication(name)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

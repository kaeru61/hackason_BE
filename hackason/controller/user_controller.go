package controller

import (
	"db/application"
	"db/model/makeupmodel"
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"log"
	"net/http"
)

func userController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		userGet(w, r)
		return
	case http.MethodPost:
		userCreate(w, r)
		return
	case http.MethodPut:
		userUpdate(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func userGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if query.Get("id") != "" {
		userId := query.Get("id")
		userInfo := application.UserGetByUserId(userId)
		if userInfo.Error.Code == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bytes, err := json.Marshal(userInfo)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	} else {
		userName := query.Get("name")
		userInfo := application.UserGetByUserName(userName)
		if userInfo.Error.Code == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bytes, err := json.Marshal(userInfo)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	var userC makeupmodel.UserCUD
	type Info struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Bio  string `json:"bio"`
	}
	var info Info

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	idA := ulid.Make()
	id := idA.String()
	userC.User.Id = id
	userC.User.Name = info.Name
	userC.User.Age = info.Age
	userC.User.Bio = info.Bio
	err := application.UserCreate(userC)
	if err.Code == 1 {
		log.Printf("fail: application.UserCreate, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err_ := json.Marshal(err)
	if err_ != nil {
		log.Printf("fail: json.Marshal, %v\n", err_)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func userUpdate(w http.ResponseWriter, r *http.Request) {
	var userU makeupmodel.UserCUD
	if err := json.NewDecoder(r.Body).Decode(&userU); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := application.UserUpdate(userU)
	if err.Code == 1 {
		log.Printf("fail: application.UserUpdate, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err_ := json.Marshal(err)
	if err_ != nil {
		log.Printf("fail: json.Marshal, %v\n", err_)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

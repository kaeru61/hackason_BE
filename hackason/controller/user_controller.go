package controller

import (
	"db/application"
	"db/model/makeupmodel"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func userController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONT_END_DOMAIN"))
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
	if query.Get("userId") != "" {
		userId := query.Get("userId")
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
		userName := query.Get("userName")
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

	if err := json.NewDecoder(r.Body).Decode(&userC); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}

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

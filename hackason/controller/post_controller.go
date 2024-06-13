package controller

import (
	"db/application"
	"db/model/makeupmodel"
	"encoding/json"
	"log"
	"net/http"
)

func postController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		postGet(w, r)
		return
	case http.MethodPost:
		postCreate(w, r)
		return
	case http.MethodDelete:
		postDelete(w, r)
		return
	case http.MethodPut:
		postUpdate(w, r)
		return
	}
}

func postGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	postId := query.Get("postId")

	postInfo := application.PostGet(postId)
	if postInfo.Error.Code == 1 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(postInfo)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func postCreate(w http.ResponseWriter, r *http.Request) {
	var postC makeupmodel.PostCUD

	if err := json.NewDecoder(r.Body).Decode(&postC); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}

	err := application.PostCreate(postC)
	if err.Code == 1 {
		log.Printf("fail: application.PostCreate, %v\n", err)
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

func postUpdate(w http.ResponseWriter, r *http.Request) {
	var postU makeupmodel.PostCUD
	if err := json.NewDecoder(r.Body).Decode(&postU); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := application.PostUpdate(postU)
	if err.Code == 1 {
		log.Printf("fail: application.PostUpdate, %v\n", err)
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

func postDelete(w http.ResponseWriter, r *http.Request) {
	var postD makeupmodel.PostCUD
	if err := json.NewDecoder(r.Body).Decode(&postD); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := application.PostDelete(postD)
	if err.Code == 1 {
		log.Printf("fail: application.PostDelete, %v\n", err)
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

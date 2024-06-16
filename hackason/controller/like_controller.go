package controller

import (
	"db/application"
	"db/model/makeupmodel"
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"log"
	"net/http"
)

func likeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		likeCreate(w, r)
		return
	case http.MethodDelete:
		likeDelete(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func likeCreate(w http.ResponseWriter, r *http.Request) {
	var likeC makeupmodel.LikeCD
	type LikeCInfo struct {
		UserId   string `json:"userId"`
		PostId   string `json:"postId"`
		CreateAt string `json:"createAt"`
	}
	var likeCInfo LikeCInfo
	if err := json.NewDecoder(r.Body).Decode(&likeCInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	likeC.UserId = likeCInfo.UserId
	likeC.PostId = likeCInfo.PostId
	likeC.CreateAt = likeCInfo.CreateAt
	likeC.Id = ulid.Make().String()

	err := application.LikeCreate(likeC)
	if err.Code == 1 {
		log.Printf("fail: application.FollowsCreate, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func likeDelete(w http.ResponseWriter, r *http.Request) {
	var likeD makeupmodel.LikeCD

	if err := json.NewDecoder(r.Body).Decode(&likeD); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := application.LikeDelete(likeD)
	if err.Code == 1 {
		log.Printf("fail: application.FollowsCreate, %v\n", err)
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

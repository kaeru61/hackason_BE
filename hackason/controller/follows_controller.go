package controller

import (
	"db/application"
	"db/model/makeupmodel"
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"log"
	"net/http"
)

func followsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		followsGet(w, r)
		return
	case http.MethodPost:
		followsCreate(w, r)
		return
	case http.MethodDelete:
		followsDelete(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func followsGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userId := query.Get("id")

	followsInfo := application.FollowsGet(userId)
	if followsInfo.Error.Code == 1 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(followsInfo)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func followsCreate(w http.ResponseWriter, r *http.Request) {
	var followsC makeupmodel.FollowsCD
	type FollowsCInfo struct {
		FollowingUId string `json:"followingUId"`
		FollowerUId  string `json:"followerUId"`
		CreateAt     string `json:"createAt"`
	}
	var followsCInfo FollowsCInfo
	if err := json.NewDecoder(r.Body).Decode(&followsCInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	followsC.FollowingUId = followsCInfo.FollowingUId
	followsC.FollowerUId = followsCInfo.FollowerUId
	followsC.CreateAt = followsCInfo.CreateAt
	followsC.Id = ulid.Make().String()

	err := application.FollowsCreate(followsC)
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

func followsDelete(w http.ResponseWriter, r *http.Request) {
	var followsD makeupmodel.FollowsCD
	if err := json.NewDecoder(r.Body).Decode(&followsD); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := application.FollowsDelete(followsD)
	if err.Code == 1 {
		log.Printf("fail: application.followsDelete, %v\n", err)
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

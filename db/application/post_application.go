package application

import (
	"db/dao/checkdao"
	"db/dao/postdao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func PostGet(postId string) makeupmodel.PostInfo {
	postInfo, err := postdao.PostGet(postId)
	if err != nil {
		log.Println("an error occurred at application/PostGet", err)
		return postInfo
	}
	return postInfo
}

func PostCreate(postC makeupmodel.PostCUD) mainmodel.Error {
	fmt.Println(postC.Post.Id)
	return postdao.PostCreate(postC.Post)
}

func PostDelete(postD makeupmodel.PostCUD) mainmodel.Error {
	posterId, err := checkdao.PostPosterId(postD.Post.Id)
	if err.Code != 0 {
		return err
	}
	if posterId != postD.Post.UserId {
		return mainmodel.Error{45, "no authority to delete post"}
	}
	fmt.Println(postD.Post.Id)
	return postdao.PostDelete(postD.Post.Id)
}

func PostUpdate(postU makeupmodel.PostCUD) mainmodel.Error {
	fmt.Println(postU.Post.Id)
	return postdao.PostUpdate(postU.Post)
}

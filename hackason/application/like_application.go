package application

import (
	"db/dao/likedao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
)

func LikeCreate(likeC makeupmodel.LikeCD) mainmodel.Error {
	fmt.Println(likeC.Like.Id)
	return likedao.LikeCreate(likeC.Like)
}

func LikeDelete(likeD makeupmodel.LikeCD) mainmodel.Error {
	fmt.Println(likeD.Like.Id)
	return likedao.LikeDelete(likeD.Like.Id)
}

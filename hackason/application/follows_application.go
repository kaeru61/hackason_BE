package application

import (
	"db/dao/followsdao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func FollowsGet(userId string) makeupmodel.FollowsInfo {
	followsInfo, err := followsdao.FollowsGet(userId)
	if err != nil {
		log.Println("An error occurred at applocation/FollowsGet", err)
		return followsInfo
	}
	return followsInfo
}

func FollowsCreate(followsC makeupmodel.FollowsCUD) mainmodel.Error {
	fmt.Println(followsC.Follows.Id)
	return followsdao.FollowsCreate(followsC.Follows)
}

func FollowsDelete(followsD makeupmodel.FollowsCUD) mainmodel.Error {
	fmt.Println(followsD.Follows.Id)
	return followsdao.FollowsDelete(followsD.Follows.Id)
}

package application

import (
	"db/dao/followsdao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
)

func FollowsCreate(followsC makeupmodel.FollowsCUD) mainmodel.Error {
	fmt.Println(followsC.Follows.Id)
	return followsdao.FollowsCreate(followsC.Follows)
}

func FollowsDelete(followsD makeupmodel.FollowsCUD) mainmodel.Error {
	fmt.Println(followsD.Follows.Id)
	return followsdao.FollowsDelete(followsD.Follows.Id)
}

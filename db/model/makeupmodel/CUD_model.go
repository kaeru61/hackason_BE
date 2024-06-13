package makeupmodel

import "db/model/mainmodel"

type PostCUD struct {
	mainmodel.Post `json:"post"`
}

type FollowsCUD struct {
	mainmodel.Follows `json:"follows"`
}

type UserCUD struct {
	mainmodel.User `json:"user"`
}

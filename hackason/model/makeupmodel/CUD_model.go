package makeupmodel

import "db/model/mainmodel"

type PostCUD struct {
	mainmodel.Post `json:"post"`
}

type FollowsCD struct {
	mainmodel.Follows `json:"follows"`
}

type UserCUD struct {
	mainmodel.User `json:"user"`
}

type LikeCD struct {
	mainmodel.Like `json:"like"`
}

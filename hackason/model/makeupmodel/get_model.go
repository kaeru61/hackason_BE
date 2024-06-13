package makeupmodel

import "db/model/mainmodel"

type PostInfo struct {
	Root            mainmodel.Post   `json:"root"`
	Replies         []mainmodel.Post `json:"replies"`
	mainmodel.Error `json:"error"`
}

type FollowsInfo struct {
	User            mainmodel.User   `json:"user"`
	Followees       []mainmodel.User `json:"followees"`
	Followers       []mainmodel.User `json:"followers"`
	mainmodel.Error `json:"error"`
}

type UserInfo struct {
	User            mainmodel.User `json:"user"`
	mainmodel.Error `json:"error"`
}

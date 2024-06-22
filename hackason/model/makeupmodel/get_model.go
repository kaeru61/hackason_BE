package makeupmodel

import "db/model/mainmodel"

type PostInfo struct {
	Root            mainmodel.Post   `json:"root"`
	User            mainmodel.User   `json:"user"`
	Replies         []mainmodel.Post `json:"replies"`
	LikedBy         []mainmodel.User `json:"likedBy"`
	mainmodel.Error `json:"error"`
}

type FollowsInfo struct {
	User            mainmodel.User   `json:"user"`
	Followings      []mainmodel.User `json:"followings"`
	Followers       []mainmodel.User `json:"followers"`
	mainmodel.Error `json:"error"`
}

type UserInfo struct {
	User            mainmodel.User   `json:"user"`
	Followings      []mainmodel.User `json:"followings"`
	Followers       []mainmodel.User `json:"followers"`
	Posts           []mainmodel.Post `json:"posts"`
	Likes           []mainmodel.Post `json:"likes"`
	mainmodel.Error `json:"error"`
}

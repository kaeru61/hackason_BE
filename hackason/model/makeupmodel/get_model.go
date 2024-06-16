package makeupmodel

import "db/model/mainmodel"

type PostInfo struct {
	Root            mainmodel.Post   `json:"root"`
	Replies         []mainmodel.Post `json:"replies"`
	mainmodel.Error `json:"error"`
}

type FollowsInfo struct {
	User            mainmodel.User   `json:"user"`
	Followings      []mainmodel.User `json:"followings"`
	Followers       []mainmodel.User `json:"followers"`
	mainmodel.Error `json:"error"`
}

type UserInfo struct {
	User            mainmodel.User `json:"user"`
	mainmodel.Error `json:"error"`
}

type LikeInfoAboutPost struct {
	Post            mainmodel.Post   `json:"post"`
	User            []mainmodel.User `json:"user"`
	mainmodel.Error `json:"error"`
}

type LikeInfoAboutUser struct {
	User            mainmodel.User   `json:"user"`
	Post            []mainmodel.Post `json:"post"`
	mainmodel.Error `json:"error"`
}

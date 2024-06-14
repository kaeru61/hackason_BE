package mainmodel

type Follows struct {
	FollowingUId string `json:"followingUId"`
	FollowerUId  string `json:"followerUId"`
	CreateAt     string `json:"createAt"`
	Id           string `json:"id"`
}

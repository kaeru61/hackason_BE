package mainmodel

type Follows struct {
	FollowingUId string `json:"followingUId"`
	FollowerUId  string `json:"followerUId"`
	CreatedAt    string `json:"createdAt"`
}

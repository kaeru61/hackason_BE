package mainmodel

type Like struct {
	Id       string `json:"id"`
	UserId   string `json:"userId"`
	PostId   string `json:"postId"`
	CreateAt string `json:"createAt"`
}

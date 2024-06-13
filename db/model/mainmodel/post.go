package mainmodel

type Post struct {
	UserId   string `json:"userId"`
	Id       string `json:"id"`
	Body     string `json:"body"`
	ParentId string `json:"parentId"`
	CreateAt string `json:"createAt"`
	Deleted  bool   `json:"deleted"`
}

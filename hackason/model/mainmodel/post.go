package mainmodel

type Post struct {
	Id       string `json:"id"`
	UserId   string `json:"userId"`
	Body     string `json:"body"`
	ParentId string `json:"parentId"`
	CreateAt string `json:"createAt"`
	Deleted  bool   `json:"deleted"`
}

package makeupmodel

import "db/model/mainmodel"

type PostInfo struct {
	Root            mainmodel.Post   `json:"root"`
	Replies         []mainmodel.Post `json:"replies"`
	mainmodel.Error `json:"error"`
}

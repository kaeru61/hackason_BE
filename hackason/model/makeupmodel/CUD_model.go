package makeupmodel

import "db/model/mainmodel"

type PostCUD struct {
	mainmodel.Post `json:"post"`
}

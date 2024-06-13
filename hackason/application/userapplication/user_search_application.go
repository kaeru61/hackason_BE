package userapplication

import (
	"db/dao"
	"db/model/mainmodel"
)

func UserSearchApplication(name string) ([]mainmodel.User, error) {
	return dao.DaoSearch(name)
}

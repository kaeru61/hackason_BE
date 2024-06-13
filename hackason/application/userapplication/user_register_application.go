package userapplication

import (
	"db/dao"
)

func UserRegisterApplication(id string, name string, age int) error {
	err := dao.DaoRegister(id, name, age)
	return err
}

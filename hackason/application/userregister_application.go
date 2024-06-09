package application

import (
	"database/sql"
	"db/dao"
)

func UserRegisterApplication(id string, name string, age int, db *sql.DB) error {
	err := dao.DaoRegister(id, name, age, db)
	return err
}

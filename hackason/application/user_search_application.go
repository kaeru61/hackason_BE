package application

import (
	"database/sql"
	"db/dao"
	"db/model"
)

func UserSearchApplication(a string, b *sql.DB) ([]model.User, error) {
	name := a
	db := b
	return dao.DaoSearch(name, db)
}

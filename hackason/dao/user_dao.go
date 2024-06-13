package dao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"log"
)

func DaoSearch(name string) ([]mainmodel.User, error) {
	rows, err := maindao.Db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var users []mainmodel.User
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func DaoRegister(id string, name string, age int) error {
	_, err := maindao.Db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", id, name, age)
	return err
}

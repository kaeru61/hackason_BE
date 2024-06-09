package dao

import (
	"database/sql"
	"db/model"
	"log"
)

func DaoSearch(name string, db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func DaoRegister(id string, name string, age int, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", id, name, age)
	return err
}

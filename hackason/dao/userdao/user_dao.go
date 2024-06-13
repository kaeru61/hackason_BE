package userdao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func UserGetUserByUserId(userId string) (makeupmodel.UserInfo, error) {
	var userInfo makeupmodel.UserInfo
	if err := userGetUserByUserId(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	return userInfo, nil
}

func UserGetUserByUserName(userName string) (makeupmodel.UserInfo, error) {
	var userInfo makeupmodel.UserInfo
	if err := userGetUserByUserName(userName, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	return userInfo, nil
}

func userGetUserByUserId(userId string, userInfo *makeupmodel.UserInfo) error {
	rows, err := maindao.Db.Query("select id, name, age, bio from user where id = ?", userId)
	if err != nil {
		log.Printf("fail: hackason.Query @userGetUserByUserId, %v\n", err)
		userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @userGetUserByUserId, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Bio, &u.Age,
		); err != nil {
			log.Printf("fail: rows.Scan @userGetUserByUserId, %v\n", err)
			userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @userGetUserByUserId, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @userGetUserByUserId, %v\n", err_)
				userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @userGetUserByUserId, %v\n", err))
				return err
			}
			return err
		}
		userInfo.User = u
	}
	return nil
}

func userGetUserByUserName(userName string, userInfo *makeupmodel.UserInfo) error {
	rows, err := maindao.Db.Query("select id, name, age, bio from user where name = ?", userName)
	if err != nil {
		log.Printf("fail: hackason.Query @userGetUserByUserName, %v\n", err)
		userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @userGetUserByUserName, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Bio, &u.Age,
		); err != nil {
			log.Printf("fail: rows.Scan @userGetUserByUserName, %v\n", err)
			userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @userGetUserByUserName, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @userGetUserByUserName, %v\n", err_)
				userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @userGetUserByUserName, %v\n", err))
				return err
			}
			return err
		}
		userInfo.User = u
	}
	return nil
}

func UserCreate(u mainmodel.User) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @user_create_dao\n", err))
	}
	rows, err := tx.Prepare("insert into user (id, name, age, bio) values(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @user_create_dao\n", err))
	}
	if _, err := rows.Exec(u.Id, u.Name, u.Age, u.Bio); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @user_create_dao\n", err))
	}
	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @user_create_dao\n", err))
	}
	log.Printf("successfully created (%+v)", u)
	return mainmodel.NilError
}

func UserUpdate(u mainmodel.User) mainmodel.Error {
	tx, err := maindao.Db.Begin()

	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @user_update_dao\n", err))
	}

	rows, err := tx.Prepare("update user set (name, bio, age) where id=?")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @user_update_dao\n", err))
	}

	if _, err := rows.Exec(u.Name, u.Bio, u.Age, u.Id); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @user_update_dao\n", err))
	}

	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @user_update_dao\n", err))
	}

	log.Printf("successfully updated post (ID: %s)", u.Id)

	return mainmodel.NilError
}

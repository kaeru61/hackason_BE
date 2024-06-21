package followsdao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func FollowsGetFollowing(userId string, userInfo *makeupmodel.UserInfo) error {
	rows, err := maindao.Db.Query(`SELECT user.id, user.name, user.age, user.bio FROM user INNER JOIN follows ON user.id = follows.followerUId WHERE followingUId = ?`, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @followsGetFollowing, %v\n", err)
		userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @messageGetReply, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Age, &u.Bio,
		); err != nil {
			log.Printf("fail: rows.Scan @followsGetFollowing, %v\n", err)
			userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @followsGetFollowing, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @followsGetFollowing, %v\n", err_)
				userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @followsGetFollowing, %v\n", err))
				return err
			}

			return err
		}
		userInfo.Followings = append(userInfo.Followings, u)
	}
	return nil
}

func FollowsGetFollower(userId string, userInfo *makeupmodel.UserInfo) error {
	rows, err := maindao.Db.Query(`SELECT user.id, user.name, user.age, user.bio FROM user INNER JOIN follows ON user.id = follows.followingUId WHERE followerUId = ?`, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @followsGetFollower, %v\n", err)
		userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @followsGetFollower, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Age, &u.Bio,
		); err != nil {
			log.Printf("fail: rows.Scan @followsGetFollower, %v\n", err)
			userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @followsGetFollower, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @followsGetFollower, %v\n", err_)
				userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @followsGetFollower, %v\n", err))
				return err
			}

			return err
		}
		userInfo.Followers = append(userInfo.Followers, u)
	}
	return nil
}

func FollowsCreate(f mainmodel.Follows) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		log.Printf("fail: hackason.Begin @follows_create_dao, %v\n", err)
	}
	rows, err := tx.Prepare(`insert into follows (followingUId, followerUId, createAt, id) values (?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @follows_create_dao\n", err))
	}
	if _, err := rows.Exec(f.FollowingUId, f.FollowerUId, f.CreateAt, f.Id); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @follows_create_dao\n", err))
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @follows_create_dao\n", err))
	}
	log.Printf("successfully created (%+v)", f)
	return mainmodel.NilError
}

func FollowsDelete(followsId string) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		log.Printf("fail: hackason.Begin @follows_delete_dao, %v\n", err)
	}
	rows, err := tx.Prepare(`delete from follows where id = ?`)
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @follows_delete_dao\n", err))
	}
	_, err = rows.Exec(followsId)
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @follows_delete_dao\n", err))
	}
	log.Printf("successfully deleted (%+v)", followsId)
	return mainmodel.NilError
}

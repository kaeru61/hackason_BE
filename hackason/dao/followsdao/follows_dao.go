package followsdao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func FollowsGet(userId string) (makeupmodel.FollowsInfo, error) {
	var followsInfo makeupmodel.FollowsInfo
	if err := followsGetUser(userId, &followsInfo); err != nil {
		log.Println("An error occured at follows_dao.go", err)
		return followsInfo, err
	}
	if err := followsGetFollowing(userId, &followsInfo); err != nil {
		log.Println("An error occured at follows_dao.go", err)
		return followsInfo, err
	}
	if err := followsGetFollower(userId, &followsInfo); err != nil {
		log.Println("An error occured at follows_dao.go", err)
		return followsInfo, err
	}
	return followsInfo, nil
}

func followsGetUser(userId string, followsInfo *makeupmodel.FollowsInfo) error {
	rows, err := maindao.Db.Query(`SELECT id, name, age, bio FROM user WHERE  id = ?`, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @followsGetUser, %v\n", err)
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Age, &u.Bio,
		); err != nil {
			log.Printf("fail: rows.Scan @followsGetUser, %v\n", err)
			followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @followGetUser, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @followsGetUser, %v\n", err_)
				followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @postGetPost, %v\n", err))
				return err_
			}
			return err
		}
		followsInfo.User = u
	}
	return nil

}

func followsGetFollowing(userId string, followsInfo *makeupmodel.FollowsInfo) error {
	rows, err := maindao.Db.Query(`SELECT id, name, age, bio FROM user INNER JOIN follows ON user.id = follows.followerUId WHERE followingUId = ?`, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @followsGetFollowing, %v\n", err)
		followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @messageGetReply, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Age, &u.Bio,
		); err != nil {
			log.Printf("fail: rows.Scan @followsGetFollowing, %v\n", err)
			followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @followsGetFollowing, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @followsGetFollowing, %v\n", err_)
				followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @followsGetFollowing, %v\n", err))
				return err
			}

			return err
		}
		followsInfo.Followings = append(followsInfo.Followings, u)
	}
	return nil
}

func followsGetFollower(userId string, followsInfo *makeupmodel.FollowsInfo) error {
	rows, err := maindao.Db.Query(`SELECT id, name, age, bio FROM user INNER JOIN follows ON user.id = follows.followingUId WHERE followerUId = ?`, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @followsGetFollower, %v\n", err)
		followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @followsGetFollower, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Age, &u.Bio,
		); err != nil {
			log.Printf("fail: rows.Scan @followsGetFollower, %v\n", err)
			followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @followsGetFollower, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @followsGetFollower, %v\n", err_)
				followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @followsGetFollower, %v\n", err))
				return err
			}

			return err
		}
		followsInfo.Followers = append(followsInfo.Followers, u)
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

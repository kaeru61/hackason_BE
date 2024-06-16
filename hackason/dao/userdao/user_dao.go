package userdao

import (
	"db/dao/followsdao"
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
	if err := followsdao.FollowsGetFollowing(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	if err := followsdao.FollowsGetFollower(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	if err := userGetPosts(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	if err := userGetLikes(userId, &userInfo); err != nil {
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
	userId := userInfo.User.Id
	if err := followsdao.FollowsGetFollowing(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	if err := followsdao.FollowsGetFollower(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	if err := userGetPosts(userId, &userInfo); err != nil {
		log.Println("An Error occured at user_dao.go", err)
		return userInfo, err
	}
	if err := userGetLikes(userId, &userInfo); err != nil {
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
			&u.Id, &u.Name, &u.Age, &u.Bio,
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
			&u.Id, &u.Name, &u.Age, &u.Bio,
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

func userGetPosts(userId string, userInfo *makeupmodel.UserInfo) error {
	rows, err := maindao.Db.Query(`SELECT id, userId, body, parentId, createAt, deleted FROM post WHERE userId = ?`, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @postGetReply, %v\n", err)
		userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @postGetReply, %v\n", err))
		return err
	}
	for rows.Next() {
		var p mainmodel.Post
		if err := rows.Scan(
			&p.Id, &p.UserId, &p.Body, &p.CreateAt, &p.ParentId, &p.Deleted,
		); err != nil {
			log.Printf("fail: rows.Scan @postGetReply, %v\n", err)
			userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @postGetReply, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @postGetReply, %v\n", err_)
				userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @postGetReply, %v\n", err))
				return err
			}

			return err
		}
		userInfo.Posts = append(userInfo.Posts, p)
	}
	return nil
}

func userGetLikes(userId string, userInfo *makeupmodel.UserInfo) error {
	rows, err := maindao.Db.Query(`SELECT id, userId, body, parentId, createAt, deleted FROM(SELECT * FROM post INNER JOIN like on post.id = like.postId where like.userId = ?) `, userId)
	if err != nil {
		log.Printf("fail: hackason.Query @postGetReply, %v\n", err)
		userInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @postGetReply, %v\n", err))
		return err
	}
	for rows.Next() {
		var p mainmodel.Post
		if err := rows.Scan(
			&p.Id, &p.UserId, &p.Body, &p.CreateAt, &p.ParentId, &p.Deleted,
		); err != nil {
			log.Printf("fail: rows.Scan @postGetReply, %v\n", err)
			userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @postGetReply, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @postGetReply, %v\n", err_)
				userInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @postGetReply, %v\n", err))
				return err
			}

			return err
		}
		userInfo.Likes = append(userInfo.Likes, p)
	}
	return nil
}

func UserCreate(u mainmodel.User) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @user_create_dao\n", err))
	}
	rows, err := tx.Prepare("insert into user (id, name, age, bio) values(?, ?, ?, ?)") //fix
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

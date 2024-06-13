package followsdao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func followsGetUser(userId string, followsInfo *makeupmodel.FollowsInfo) error {
	rows, err := maindao.Db.Query(`SELECT * FROM user WHERE  id = ?`, userId)
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
	rows, err := maindao.Db.Query(`SELECT * FROM follows WHERE followingUId = ?`, userId)
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
			followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @postGetReply, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @followsGetFollowing, %v\n", err_)
				followsInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @followsGetFollowing, %v\n", err))
				return err
			}

			return err
		}
		postInfo.Replies = append(postInfo.Replies, p)
	}
	return nil
}

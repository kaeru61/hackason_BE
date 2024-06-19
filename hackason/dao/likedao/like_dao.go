package likedao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"fmt"
	"log"
)

func LikeCreate(l mainmodel.Like) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @like_create_dao\n", err))
	}
	rows, err := tx.Prepare("insert into `like` (id, userId, postId, createAt) values(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @like_create_dao\n", err))
	}
	if _, err := rows.Exec(l.Id, l.UserId, l.PostId, l.CreateAt); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @like_create_dao\n", err))
	}
	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @like_create_dao\n", err))
	}
	log.Printf("successfully created (%+v)", l)
	return mainmodel.NilError
}

func LikeDelete(likeId string) mainmodel.Error {
	tx, err := maindao.Db.Begin()

	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @like_delete_dao\n", err))
	}

	rows, err := tx.Prepare("delete from `like` where id = ?")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @like_delete_dao\n", err))
	}

	if _, err := rows.Exec(likeId); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @like_delete_dao\n", err))
	}
	log.Printf("successfully deleted relevant replies (post ID: %s)", likeId)

	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @like_delete_dao\n", err))
	}

	return mainmodel.NilError
}

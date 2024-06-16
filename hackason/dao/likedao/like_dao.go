package likedao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

//likeのみの情報はいらない

func LikeAboutPostGet(postId string) (makeupmodel.LikeInfoAboutPost, error) {
	var likeInfoPost makeupmodel.LikeInfoAboutPost
	err := likeAboutPostGetPost(postId, &likeInfoPost)
	if err != nil {
		log.Println("An Error occured at like_dao.go", err)
		return likeInfoPost, err
	}
	_err := likeAboutPostGetUser(postId, &likeInfoPost)
	if _err != nil {
		log.Println("An Error occured at like_dao.go", err)
		return likeInfoPost, _err
	}
	return likeInfoPost, nil
}

func likeAboutPostGetPost(postId string, likeInfoPost *makeupmodel.LikeInfoAboutPost) error {
	rows, err := maindao.Db.Query(`SELECT id, userId, body, parentId, createAt, deleted  FROM post WHERE id = ?`, postId)
	if err != nil {
		log.Printf("fail: hackason.Query @postGetPost, %v\n", err)
	}
	for rows.Next() {
		var p mainmodel.Post
		if err := rows.Scan(
			&p.Id, &p.UserId, &p.Body, &p.ParentId, &p.CreateAt, &p.Deleted,
		); err != nil {
			log.Printf("fail: rows.Scan @postGetPost, %v\n", err)
			likeInfoPost.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @postGetPost, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @postGetPost, %v\n", err_)
				likeInfoPost.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @postGetPost, %v\n", err))
				return err
			}
			return err
		}
		likeInfoPost.Post = p
	}
	return nil
}

func likeAboutPostGetUser(postId string, likeInfoPost *makeupmodel.LikeInfoAboutPost) error {
	rows, err := maindao.Db.Query("select id, name, age, bio from(select * from user inner join like on user.id = like.userId where like.postId = ?) ", postId)
	if err != nil {
		log.Printf("fail: hackason.Query @likeAboutPostGetUser, %v\n", err)
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Bio); err != nil {
			log.Printf("fail: hackason.Query @likeAboutPostGetUser, %v\n", err)
			likeInfoPost.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @LikeAboutPostGetUser, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @LikeAboutPostGetUser, %v\n", err_)
				likeInfoPost.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @LikeAboutPostGetUser, %v\n", err))
				return err
			}
			return err
		}
		likeInfoPost.User = append(likeInfoPost.User, u)
	}
	return nil
}

func likeAboutUserGetUser(userId string, likeInfoUser *makeupmodel.LikeInfoAboutUser) error {
	rows, err := maindao.Db.Query("select id, name, age, bio from user where id = ?", userId)
	if err != nil {
		log.Printf("fail: hackason.Query @LikeAboutUserGetUser, %v\n", err)
		likeInfoUser.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @LikeAboutUserGetUser, %v\n", err))
		return err
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(
			&u.Id, &u.Name, &u.Age, &u.Bio,
		); err != nil {
			log.Printf("fail: rows.Scan @LikeAboutUserGetUser, %v\n", err)
			likeInfoUser.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @LikeAboutUserGetUser, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @LikeAboutUserGetUser, %v\n", err_)
				likeInfoUser.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @LikeAboutUserGetUser, %v\n", err))
				return err
			}
			return err
		}
		likeInfoUser.User = u
	}
	return nil
}

func likeAboutUserGetPost(userId string, likeInfoUser *makeupmodel.LikeInfoAboutUser) error {
	rows, err := maindao.Db.Query("select id, userId, body, parentId, createAt, deleted from(select * from post inner join like on post.id = like.postId where like.userId = ?) ", userId)
	if err != nil {
		log.Printf("fail: hackason.Query @likeAboutUserGetPost, %v\n", err)
	}
	for rows.Next() {
		var p mainmodel.Post
		if err := rows.Scan(&p.Id, &p.UserId, &p.Body, &p.ParentId, &p.CreateAt, &p.Deleted); err != nil {
			log.Printf("fail: hackason.Query @likeAboutPostGetUser, %v\n", err)
			likeInfoUser.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @likeAboutUserGetPost, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @likeAboutUserGetPost, %v\n", err_)
				likeInfoUser.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @likeAboutUserGetPost, %v\n", err))
				return err
			}
			return err
		}
		likeInfoUser.Post = append(likeInfoUser.Post, p)
	}
	return nil
}

func LikeCreate(l mainmodel.Like) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @like_create_dao\n", err))
	}
	rows, err := tx.Prepare("insert into like (id, userId, postId, createAt) values(?, ?, ?, ?)")
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

	rows, err := tx.Prepare("delete from like where id = ?")
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

package postdao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func PostGetAllPost() ([]makeupmodel.PostInfo, error) {
	var posts []makeupmodel.PostInfo
	rows, err := maindao.Db.Query(`SELECT id FROM post WHERE deleted = false ORDER BY createAt DESC`)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {
			return posts, err
			if err := rows.Close(); err != nil {
				return posts, err
			}
		}
		post, _ := PostGet(p)
		posts = append(posts, post)
	}
	return posts, nil
}

func PostGet(postId string) (makeupmodel.PostInfo, error) {
	var postInfo makeupmodel.PostInfo
	if err := postGetPost(postId, &postInfo); err != nil {
		log.Println("An Error occured at post_dao.go", err)
		return postInfo, err
	}
	if err := postGetReply(postId, &postInfo); err != nil {
		log.Println("An Error occured at post_dao.go", err)
		return postInfo, err
	}
	if err := postGetLikedBy(postId, &postInfo); err != nil {
		log.Println("An Error occured at post_dao.go", err)
		return postInfo, err
	}
	return postInfo, nil
}

func postGetPost(postId string, postInfo *makeupmodel.PostInfo) error {
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
			postInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @postGetPost, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @postGetPost, %v\n", err_)
				postInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @postGetPost, %v\n", err))
				return err
			}
			return err
		}
		postInfo.Root = p
	}
	return nil

}

func postGetReply(postId string, postInfo *makeupmodel.PostInfo) error {
	rows, err := maindao.Db.Query(`SELECT id, userId, body, parentId, createAt, deleted FROM post WHERE parentId = ?`, postId)
	if err != nil {
		log.Printf("fail: hackason.Query @postGetReply, %v\n", err)
		postInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @postGetReply, %v\n", err))
		return err
	}
	for rows.Next() {
		var p mainmodel.Post
		if err := rows.Scan(
			&p.Id, &p.UserId, &p.Body, &p.CreateAt, &p.ParentId, &p.Deleted,
		); err != nil {
			log.Printf("fail: rows.Scan @postGetReply, %v\n", err)
			postInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Scan @postGetReply, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @postGetReply, %v\n", err_)
				postInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @postGetReply, %v\n", err))
				return err
			}

			return err
		}
		postInfo.Replies = append(postInfo.Replies, p)
	}
	return nil
}

func postGetLikedBy(postId string, postInfo *makeupmodel.PostInfo) error {
	rows, err := maindao.Db.Query("SELECT user.id, user.name, user.age, user.bio FROM user INNER JOIN `like` ON user.id = `like`.userId WHERE `like`.postId = ?", postId)
	if err != nil {
		log.Printf("fail: hackason.Query @likeAboutPostGetUser, %v\n", err)
	}
	for rows.Next() {
		var u mainmodel.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Bio); err != nil {
			log.Printf("fail: hackason.Query @likeAboutPostGetUser, %v\n", err)
			postInfo.Error.UpdateError(1, fmt.Sprintf("fail: hackason.Query @LikeAboutPostGetUser, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close @LikeAboutPostGetUser, %v\n", err_)
				postInfo.Error.UpdateError(1, fmt.Sprintf("fail: rows.Close @LikeAboutPostGetUser, %v\n", err))
				return err
			}
			return err
		}
		postInfo.LikedBy = append(postInfo.LikedBy, u)
	}
	return nil
}

func PostCreate(p mainmodel.Post) mainmodel.Error {
	tx, err := maindao.Db.Begin()
	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @post_create_dao\n", err))
	}
	rows, err := tx.Prepare("insert into post (id, userId, body, createAt, parentId, deleted) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @post_create_dao\n", err))
	}
	if _, err := rows.Exec(p.Id, p.UserId, p.Body, p.CreateAt, p.ParentId, false); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @post_create_dao\n", err))
	}
	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @post_create_dao\n", err))
	}
	log.Printf("successfully created (%+v)", p)
	return mainmodel.NilError
}

func PostDelete(id string) mainmodel.Error {
	tx, err := maindao.Db.Begin()

	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @post_delete_dao\n", err))
	}

	rows, err := tx.Prepare("update post set deleted=? where id=?")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @post_delete_dao\n", err))
	}

	if _, err := rows.Exec(true, id); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @post_delete_dao\n", err))
	}

	log.Printf("successfully deleted post (ID: %s)", id)

	// ------------

	rows, err = tx.Prepare("UPDATE post SET deleted=? WHERE parentId=?")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @post_delete_dao\n", err))
	}

	if _, err := rows.Exec(true, id); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @post_delete_dao\n", err))
	}
	log.Printf("successfully deleted relevant replies (post ID: %s)", id)

	// ------------

	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @post_delete_dao\n", err))
	}

	return mainmodel.NilError
}

func PostUpdate(p mainmodel.Post) mainmodel.Error {
	tx, err := maindao.Db.Begin()

	if err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Begin, %v @post_update_dao\n", err))
	}

	rows, err := tx.Prepare("update post set body=? where id=?")
	if err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Prepare, %v @post_update_dao\n", err))
	}

	if _, err := rows.Exec(p.Body, p.Id); err != nil {
		tx.Rollback()
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Exec, %v @post_update_dao\n", err))
	}

	if err := tx.Commit(); err != nil {
		return mainmodel.MakeError(1, fmt.Sprintf("fail: tx.Commit, %v @post_update_dao\n", err))
	}

	log.Printf("successfully updated post (ID: %s)", p.Id)

	return mainmodel.NilError
}

package checkdao

import (
	"db/dao/maindao"
	"db/model/mainmodel"
	"fmt"
	"log"
)

func PostPosterId(postId string) (string, mainmodel.Error) {
	rows, err := maindao.Db.Query("select userId from post where postId=?", postId)

	if err != nil {
		log.Printf("fail: hackason.Query, %v\n", err)
		return "", mainmodel.MakeError(1, fmt.Sprintf("fail: hackason.Query, %v\n", err))
	}

	var posterId string

	for rows.Next() {
		if err := rows.Scan(&posterId); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			err := mainmodel.MakeError(1, fmt.Sprintf("fail: rows.Scan, %v\n", err))

			if err_ := rows.Close(); err_ != nil {
				log.Printf("fail: rows.Close, %v\n", err_)
				err.UpdateError(1, fmt.Sprintf("fail: rows.Close, %v\n", err))
				return "", err
			}

			return "", err
		}
	}
	return posterId, mainmodel.NilError
}

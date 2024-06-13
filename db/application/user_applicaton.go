package application

import (
	"db/dao/userdao"
	"db/model/mainmodel"
	"db/model/makeupmodel"
	"fmt"
	"log"
)

func UserGetByUserId(userId string) makeupmodel.UserInfo {
	userInfo, err := userdao.UserGetUserByUserId(userId)
	if err != nil {
		log.Println("An error occurred at application/user_application", err)
		return userInfo
	}
	return userInfo
}

func UserGetByUserName(userName string) makeupmodel.UserInfo {
	userInfo, err := userdao.UserGetUserByUserName(userName)
	if err != nil {
		log.Println("An error occurred at application/user_application", err)
		return userInfo
	}
	return userInfo
}

func UserCreate(userC makeupmodel.UserCUD) mainmodel.Error {
	fmt.Println(userC.User.Id)
	return userdao.UserCreate(userC.User)
}

func UserUpdate(userU makeupmodel.UserCUD) mainmodel.Error {
	fmt.Println(userU.User.Id)
	return userdao.UserUpdate(userU.User)
}

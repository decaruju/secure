package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"secure/model"
)

func CreateUser(username string, password string) (*model.User, error) {
	db, err := db()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user, err := FindUserByUsername(username)
	if !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	user = &model.User{Username: username}
	user.CreatePassword(password)
	db.Create(user)
	return user, nil
}

func FindUserByUsername(username string) (*model.User, error) {
	db, err := db()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user model.User
	err = db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func FindUserByKey(key string) (*model.User, error) {
	db, err := db()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var apikey model.Apikey
	err = db.Where("key = ?", key).Preload("User").First(&apikey).Error
	fmt.Println(apikey)
	fmt.Println(apikey.UserID)
	fmt.Println(apikey.User)
	return &apikey.User, err
}

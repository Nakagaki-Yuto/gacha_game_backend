package puser

import (

	"go_practice_mvc/database"
)

type User struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

var db = database.GetDB()

func Create(name string, token string) error {
	return db.Create(&User{
		Name:  name,
		Token: token,
	}).Error
}

func Get(token string, user User) error {
	return db.Where("token = ?", token).Find(&user).Error
}

func Update(name string, token string) error {
	return db.Model(User{}).Where("token = ?", token).Update(&User{
		Name: name,
	}).Error
}

func GetID(token string, user *User) error {
	return db.Where("token = ?", token).Find(&user).Error
}
package puser

import (
	"fmt"

	"go_practice_mvc/database"
)

type User struct {
	ID int `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}


func Create(name string, token string) error {
	db := database.GetDB()
	return db.Create(&User{
		Name:  name,
		Token: token,
	}).Error
}

func Get(token string) (User, error) {
	db := database.GetDB()
	u := User{}
	error := db.Where("token = ?", token).Find(&u).Error

	if error != nil {
		fmt.Println(error)
	}

	return u, error
}

func Update(name string, token string) error {
	db := database.GetDB()

	return db.Model(&User{}).Where("token = ?", token).Update(&User{
		Name: name,
	}).Error
}

func GetID(token string) (User, error) {
	db := database.GetDB()
	ui := User{}
	error := db.Where("token = ?", token).Find(&ui).Error

	if error != nil {
		fmt.Println(error)
	}

	return ui, error
}
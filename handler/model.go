package handler

import (
	"fmt"

)

type User struct {
	ID int `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}


type Character struct {
	ID   string `json:"characterID"`
	Name string `json:"name"`
}

func (db *DB) GetCharacter(characterID string) (Character, error) {
	c := Character{}
	error := db.Where("id = ?", characterID).Find(&c).Error

	if error != nil {
		fmt.Println(error)
	}

	return c, error
}



type GachaRate struct {
	Rate        int    `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []*GachaRate

func (db *DB) GetGachaRate() (GachaRates, error) {
	gr := GachaRates{}
	error := db.Find(&gr).Error

	if error != nil {
		fmt.Println(error)
	}

	return gr, error

}



type UserCharacter struct {
	UserID int `json:"userID"`
	CharacterID string `json:"characterID"`
}

type UserCharacters []UserCharacter

func (db *DB) GetUserCharacter(userID int) (UserCharacters, error) {
	uc := UserCharacters{}
	error := db.Where("user_id = ?", userID).Find(&uc).Error

	if error != nil {
		fmt.Println(error)
	}
	return uc, error
}

func (db *DB) CreateUserCharacter(userID int, characterID string) error {
	return db.Create(&UserCharacter{
		UserID: userID,
		CharacterID: characterID,
	}).Error
}


type User struct {
	ID int `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}


func (db *DB) CreateUser(name string, token string) error {
	return db.Create(&User{
		Name:  name,
		Token: token,
	}).Error
}

func (db *DB) GetUser(token string) (User, error) {
	u := User{}
	error := db.Where("token = ?", token).Find(&u).Error

	if error != nil {
		fmt.Println(error)
	}

	return u, error
}

func (db *DB) UpdateUser(name string, token string) error {

	return db.Model(&User{}).Where("token = ?", token).Update(&User{
		Name: name,
	}).Error
}

func (db *DB) GetUserID(token string) (User, error) {

	ui := User{}
	error := db.Where("token = ?", token).Find(&ui).Error

	if error != nil {
		fmt.Println(error)
	}

	return ui, error
}
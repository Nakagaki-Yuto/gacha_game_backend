package model

import (
	"github.com/jinzhu/gorm"
)

// user model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Users []User

func CreateUser(db *gorm.DB, name string, token string) error {
	return db.Create(&User{
		Name:  name,
		Token: token,
	}).Error
}

func GetUser(db *gorm.DB, token string) (User, error) {
	u := User{}
	err := db.Where("token = ?", token).Find(&u).Error

	if err != nil {
		return User{}, err
	}

	return u, nil

}

func UpdateUser(db *gorm.DB, name string, token string) error {

	return db.Model(&User{}).Where("token = ?", token).Update(&User{
		Name: name,
	}).Error
}

func GetUserID(db *gorm.DB, token string) (User, error) {

	u := User{}
	err := db.Where("token = ?", token).Find(&u).Error

	if err != nil {
		return User{}, err
	}

	return u, nil
}

func GetAllUsers(db *gorm.DB) (Users, error) {
	u := Users{}
	err := db.Find(&u).Error

	if err != nil {
		return Users{}, err
	}

	return u, nil
}

func GetUserName(db *gorm.DB, userID int) (User, error) {

	u := User{}
	err := db.Where("id = ?", userID).Find(&u).Error

	if err != nil {
		return User{}, err
	}

	return u, nil
}

// character model

type Character struct {
	ID    string `json:"characterID"`
	Name  string `json:"name"`
	Power int    `json:"power"`
}

func GetCharacter(db *gorm.DB, characterID string) (Character, error) {
	c := Character{}
	err := db.Where("id = ?", characterID).Find(&c).Error

	if err != nil {
		return Character{}, err
	}

	return c, nil
}


// user_character model

type UserCharacter struct {
	UserID      int    `json:"userID"`
	CharacterID string `json:"characterID"`
}

type UserCharacters []UserCharacter

func GetUserCharacters(db *gorm.DB, userID int) (UserCharacters, error) {
	uc := UserCharacters{}
	err := db.Where("user_id = ?", userID).Find(&uc).Error

	if err != nil {
		return UserCharacters{}, err
	}
	return uc, nil
}

func CreateUserCharacter(db *gorm.DB, userID int, characterID string) error {
	return db.Create(&UserCharacter{
		UserID:      userID,
		CharacterID: characterID,
	}).Error
}


// gacha model

type GachaRate struct {
	Rate        int    `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []GachaRate

func GetGachaRate(db *gorm.DB) (GachaRates, error) {
	gr := GachaRates{}
	err := db.Find(&gr).Error

	if err != nil {
		return GachaRates{}, err
	}

	return gr, nil
}

package pusercharacter

import (

	"go_practice_mvc/database"
)

type UserCharacter struct {
	UserID int `json:"userID"`
	CharacterID string `json:"characterID"`
}

type UserCharacters []UserCharacter

var db = database.GetDB()

func Get(userID int, userCharacters *UserCharacters) error {
	return db.Where("user_id = ?", userID).Find(&userCharacters).Error
}

func Create(userID int, characterID string) error {
	return db.Create(&UserCharacter{
		UserID: userID,
		CharacterID: characterID,
	}).Error
}
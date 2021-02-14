package pusercharacter

import (
	"fmt"

	"go_practice_mvc/database"
)

type UserCharacter struct {
	UserID int `json:"userID"`
	CharacterID string `json:"characterID"`
}

type UserCharacters []UserCharacter

func Get(userID int) (UserCharacters, error) {
	db := database.GetDB()
	uc := UserCharacters{}
	error := db.Where("user_id = ?", userID).Find(&uc).Error

	if error != nil {
		fmt.Println(error)
	}
	return uc, error
}

func Create(userID int, characterID string) error {
	db := database.GetDB()
	return db.Create(&UserCharacter{
		UserID: userID,
		CharacterID: characterID,
	}).Error
}
package pcharacter

import (

	"go_practice_mvc/database"
)

var db = database.GetDB()

type Character struct {
	ID string `json:"characterID"`
	Name string `json:"name"`
}

func Get(characterID int, character *Character) error {
	return db.Where("id = ?", characterID).Find(&character).Error

}
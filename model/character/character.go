package pcharacter

import (
	"fmt"
	"go_practice_mvc/database"
)

type Character struct {
	ID   string `json:"characterID"`
	Name string `json:"name"`
}

func Get(characterID string) (Character, error) {
	db := database.GetDB()
	c := Character{}
	error := db.Where("id = ?", characterID).Find(&c).Error

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("キャラクター情報を取得しました")
	}

	return c, error
}

package user

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"go_practice_mvc/database"
)

type Character struct {
	ID   string `json:"characterID"`
	Name string `json:"name"`
}

func Get(characterID string) (Character, error) {
	c := Character{}
	error := db.Where("id = ?", characterID).Find(&c).Error

	if error != nil {
		fmt.Println(error)
	}

	return c, error
}

package pgacharate

import (

	"go_practice_mvc/database"
)

type GachaRate struct {
	Rate int `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []*GachaRate

var db = database.GetDB()

func Get(gachaRates *GachaRates) error {
	return db.Find(gachaRates).Error
}
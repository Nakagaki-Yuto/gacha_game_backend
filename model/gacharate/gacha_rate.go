package pgacharate

import (
	"fmt"

	"go_practice_mvc/database"
)

type GachaRate struct {
	Rate int `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []*GachaRate

func Get() (GachaRates, error) {
	db := database.GetDB()
	gr := GachaRates{}
	error := db.Find(&gr).Error

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("ガチャ情報を取得しました")
	}

	return gr, error

}
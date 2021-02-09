package gacha

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var db *gorm.DB

type User struct {
	ID int `json:"id"`
	Token string `json:"token"`
}
type GachaDrawRequest struct {
	Times int `json:"times"`
}

type Result struct {
	Result Characters `json:"results"`
}

type UserCharacter struct {
	UserID int `json:"userID"`
	CharacterID string `json:"characterID"`
}

type UserCharacters []UserCharacter

type Character struct {
	ID string `json:"characterID"`
	Name string `json:"name"`
}

type Characters []Character

type GachaRate struct {
	Rate int `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []*GachaRate

// ガチャ実行
func DrawGacha(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user := User{}
	db.Where("token = ?", token).Find(&user)
	userID := user.ID
	u := new(GachaDrawRequest)
    if err := c.Bind(u); err != nil {
        return err
    }
	times := u.Times
	fmt.Println(times)
	var characters Characters

	for i := 0; i < times; i++ {
		targetCharacter := Gacha()
		db.Create(&UserCharacter{
			UserID: userID,
			CharacterID: targetCharacter,
		})

		character := Character{}
		db.Where("id = ?", targetCharacter).Find(&character)
		characters = append(characters, character)
	}

	return c.JSON(http.StatusOK, Result{Result: characters})
}

// ガチャを引く
func Gacha() string {
	gachaRates := GachaRates{}
	db.Find(&gachaRates)
	kind := len(gachaRates)
	rates := make([]int, kind)
	maxCnt := 0

	for i := 0; i < kind; i ++ {
		targetRate := gachaRates[i].Rate
		maxCnt += targetRate
		rates[i] = targetRate
	} 

	parameter := make([]string, maxCnt)
	targetCnt:= 0
	for i := 0; i < kind; i ++ {
		for j := targetCnt; j < targetCnt + rates[i]; j ++ {
			parameter[j] = gachaRates[i].CharacterID
		}
		targetCnt += rates[i]
	} 

	rand.Seed(time.Now().UnixNano())	

	return parameter[rand.Intn(maxCnt)]
	
}
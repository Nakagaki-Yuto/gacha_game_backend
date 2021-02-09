package character

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var db *gorm.DB

type User struct {
	ID int `json:"id"`
	Token string `json:"token"`
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

type Result struct {
	Result UserCharacterLists `json:"characters"`
}

type UserCharacterList struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID string `json:"characterID"`
	Name string `json:"name"` 
}

type UserCharacterLists []UserCharacterList

// ユーザ所持キャラクター一覧取得
func GetCharacterList(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user := User{}
	db.Where("token = ?", token).Find(&user)
	userID := user.ID
	userCharacters := UserCharacters{}

	db.Where("user_id = ?", userID).Find(&userCharacters)
	
	var userCharacterLists UserCharacterLists

	for i := 0; i < len(userCharacters); i ++ {
		targetCharacterID := userCharacters[i].CharacterID
		userCharacterList := UserCharacterList{}
		userCharacterList.CharacterID = targetCharacterID
		character := Character{}
		db.Where("id = ?", targetCharacterID).Find(&character)
		userCharacterList.Name = character.Name
		userCharacterList.UserCharacterID = fmt.Sprintf("%d_%d", userID, i+1)  // "userID_所持順"のフォーマット
		userCharacterLists = append(userCharacterLists, userCharacterList)
	}

	return c.JSON(http.StatusOK, Result{Result: userCharacterLists})
}
package character

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"go_practice_mvc/model/user"
	"go_practice_mvc/model/usercharacter"
	"go_practice_mvc/model/character"
)

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
	user, error := puser.GetID(token)
	
	if error != nil {
		fmt.Println(error)
		return error
	} else {
		fmt.Println("ユーザ情報を取得しました")
	}

	userID := user.ID
	userCharacters, error := pusercharacter.Get(userID)

	if error != nil {
		fmt.Println(error)
		return error
	} else {
		fmt.Println("ユーザ‗キャラクター情報を取得しました")
	}
	
	var userCharacterLists UserCharacterLists

	for i := 0; i < len(userCharacters); i ++ {
		characterID := userCharacters[i].CharacterID
		userCharacterList := UserCharacterList{}
		userCharacterList.CharacterID = characterID
		character, error := pcharacter.Get(characterID)

		if error != nil {
			fmt.Println(error)
			return error
		} else {
			fmt.Println("キャラクター情報を取得しました")
		}
		
		userCharacterList.Name = character.Name
		userCharacterList.UserCharacterID = fmt.Sprintf("%d_%d", userID, i+1)  // "userID_所持順"のフォーマット
		userCharacterLists = append(userCharacterLists, userCharacterList)
	}

	return c.JSON(http.StatusOK, Result{Result: userCharacterLists})
}
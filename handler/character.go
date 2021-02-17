package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

)



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
func (h *Handler) GetCharacterList(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user, error := h.db.GetUserID(token)
	
	if error != nil {
		fmt.Println(error)
		return error
	}

	userID := user.ID
	userCharacters, error := h.db.GetUserCharacter(userID)

	if error != nil {
		fmt.Println(error)
		return error
	}
	
	var userCharacterLists UserCharacterLists

	for i := 0; i < len(userCharacters); i ++ {
		characterID := userCharacters[i].CharacterID
		userCharacterList := UserCharacterList{}
		userCharacterList.CharacterID = characterID
		character, error := h.db.GetCharacter(characterID)

		if error != nil {
			fmt.Println(error)
			return error
		}
		
		userCharacterList.Name = character.Name
		userCharacterList.UserCharacterID = fmt.Sprintf("%d_%d", userID, i+1)  // "userID_所持順"のフォーマット
		userCharacterLists = append(userCharacterLists, userCharacterList)
	}

	fmt.Println("ユーザ-キャラクター一覧を取得しました")
	return c.JSON(http.StatusOK, Result{Result: userCharacterLists})
}
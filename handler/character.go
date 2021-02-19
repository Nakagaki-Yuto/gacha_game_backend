package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"go_practice_mvc/model"
)

type CharacterListResponse struct {
	Characters UserCharacters `json:"characters"`
}

type UserCharacter struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type UserCharacters []UserCharacter

// ユーザ所持キャラクター一覧取得
func (h *Handler) GetCharacterList(c echo.Context) error {

	t := c.Request().Header.Get("x-token")
	u, err := model.GetUserID(h.db, t)

	if err != nil {
		fmt.Println(err)
		return err
	}

	uI := u.ID
	userCharas, err := model.GetUserCharacters(h.db, uI)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var userCharacters UserCharacters

	for i := 0; i < len(userCharas); i++ {
		cI := userCharas[i].CharacterID
		uc := UserCharacter{}
		uc.CharacterID = cI
		chara, err := model.GetCharacter(h.db, cI)

		if err != nil {
			fmt.Println(err)
			return err
		}

		uc.Name = chara.Name
		uc.UserCharacterID = fmt.Sprintf("%d_%d", uI, i+1) // "userID_所持順"のフォーマット

		userCharacters = append(userCharacters, uc)
	}

	fmt.Println("ユーザ-キャラクター一覧を取得しました")

	return c.JSON(http.StatusOK, CharacterListResponse{Characters: userCharacters})
}
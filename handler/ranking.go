package handler

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"

	"go_practice_mvc/model"
)

/*Responses
{
  "ranking": [
    {
      "name": "string",
      "maxPower": "string"
    }
  ]
}*/

type GetRankingResponse struct {
	Ranking UserMaxPowers `json:"ranking"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Users []User

type UserMaxPower struct {
	Name     string `json:"name"`
	MaxPower int `json:"maxPower"`
}

type UserMaxPowers []UserMaxPower

// ユーザ所持キャラクター一覧取得
func (h *Handler) GetRanking(c echo.Context) error {

	allUsers, err := model.GetAllUsers(h.db)

	if err != nil {
		return ErrorHandler(&err, c)
	}

	var userMaxPowers UserMaxPowers

	for i := 0; i < len(allUsers); i++ {
		ump := UserMaxPower{}
		uI := allUsers[i].ID
		u, err := model.GetUserName(h.db, uI)

		if err != nil {
			return ErrorHandler(&err, c)
		}

		ump.Name = u.Name
		userCharas, err := model. GetUserCharacters(h.db, uI)

		if err != nil {
			return ErrorHandler(&err, c)
		}

		mp := 0
		for j:=0; j < len(userCharas); j ++ {
			chara, err := model.GetCharacter(h.db, userCharas[j].CharacterID)

			if err != nil {
				return ErrorHandler(&err, c)
			}

			if chara.Power > mp {
				mp = chara.Power 
			}
		}

		ump.MaxPower = mp
		userMaxPowers = append(userMaxPowers, ump)

	}

	sort.SliceStable(userMaxPowers, func(i, j int) bool { return userMaxPowers[i].MaxPower > userMaxPowers[j].MaxPower })

	return c.JSON(http.StatusOK, GetRankingResponse{Ranking: userMaxPowers})
}

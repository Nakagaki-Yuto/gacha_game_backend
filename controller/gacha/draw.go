package gacha

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"

	pcharacter "go_practice_mvc/model/character"
	pgacharate "go_practice_mvc/model/gacharate"
	puser "go_practice_mvc/model/user"
	pusercharacter "go_practice_mvc/model/usercharacter"
)

type User struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}
type GachaDrawRequest struct {
	Times int `json:"times"`
}

type Result struct {
	Result Characters `json:"results"`
}

type UserCharacter struct {
	UserID      int    `json:"userID"`
	CharacterID string `json:"characterID"`
}

type UserCharacters []UserCharacter

type Character struct {
	ID   string `json:"characterID"`
	Name string `json:"name"`
}

type Characters []Character

type GachaRate struct {
	Rate        int    `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []GachaRate

// ガチャ実行
func DrawGacha(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user, error := puser.GetID(token)

	if error != nil {
		fmt.Println(error)
		return error
	} else {
		fmt.Println("ユーザ情報を取得しました")
	}

	userID := user.ID
	u := new(GachaDrawRequest)

	if err := c.Bind(u); err != nil {
		return err
	}

	times := u.Times

	var characters Characters

	for i := 0; i < times; i++ {
		characterID, err := Gacha()

		if err != nil {
			fmt.Println(err)
			return err
		} else {
			fmt.Println("ガチャを引きました")
		}

		error = pusercharacter.Create(userID, characterID)

		if error != nil {
			fmt.Println(error)
			return error
		} else {
			fmt.Println("ユーザー_キャラクター情報を作成しました")
		}

		character, error := pcharacter.Get(characterID)

		if error != nil {
			fmt.Println(error)
			return error
		} else {
			fmt.Println("キャラクター情報を取得しました")
		}
		c := Character{
			ID : character.ID,
			Name : character.Name,
		}

		characters = append(characters, c)
	}

	return c.JSON(http.StatusOK, Result{Result: characters})
}

// ガチャを引く
func Gacha() (string, error) {
	
	gachaRates, error := pgacharate.Get()

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("ガチャ情報を取得しました")
	}

	kind := len(gachaRates)
	rates := make([]int, kind)
	maxCnt := 0

	for i := 0; i < kind; i++ {
		targetRate := gachaRates[i].Rate
		maxCnt += targetRate
		rates[i] = targetRate
	}

	parameter := make([]string, maxCnt)
	targetCnt := 0
	for i := 0; i < kind; i++ {
		for j := targetCnt; j < targetCnt+rates[i]; j++ {
			parameter[j] = gachaRates[i].CharacterID
		}
		targetCnt += rates[i]
	}

	rand.Seed(time.Now().UnixNano())

	return parameter[rand.Intn(maxCnt)], error

}

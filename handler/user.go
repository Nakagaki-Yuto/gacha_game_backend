package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)



type UserCreateResponse struct {
	Token string `json:"token"`
}

// ユーザアカウント認証情報作成
func (h *Handler) CreateUser(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	name := u.Name
	token := CreateToken(name)
	error := h.db.CreateUser(name, token)

	if error != nil {
		fmt.Println(error)
		return error
	}

	fmt.Println("ユーザアカウント認証情報を作成しました")

	return c.JSON(http.StatusOK, UserCreateResponse{Token: token})
}

// ト-クンの作成
func CreateToken(name string) string {

	// 鍵となる文字列
	secret := "secret"

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"iss":  "init",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

type UserGetResponse struct {
	Name string `json:"name"`
}

// ユーザ情報取得
func (h *Handler) GetUser(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user, error := h.db.GetUser(token)

	if error != nil {
		fmt.Println(error)
		return error
	}

	fmt.Println("ユーザ情報を取得しました")

	return c.JSON(http.StatusOK, UserGetResponse{Name: user.Name})
}



// ユーザ情報更新
func (h *Handler) UpdateUser(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	name := u.Name
	token := c.Request().Header.Get("x-token")
	error := h.db.UpdateUser(name, token)

	if error != nil {
		fmt.Println(error)
		return error
	}

	fmt.Println("ユーザ情報を更新しました")

	return c.JSON(http.StatusOK, "name change succeed")
}

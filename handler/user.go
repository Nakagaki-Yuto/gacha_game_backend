package handler

import (
	"fmt"
	"log"
	"net/http"
	// "go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"go_practice_mvc/model"
)

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

// ユーザアカウント認証情報作成
func (h *Handler) CreateUser(c echo.Context) error {

	req := new(UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	n := req.Name
	t := CreateToken(n)
	err := model.CreateUser(h.db, n, t)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("ユーザアカウント認証情報を作成しました")

	return c.JSON(http.StatusOK, UserCreateResponse{Token: t})
}

// ト-クンの作成
func CreateToken(name string) string {

	// 鍵となる文字列
	secret := "secret"

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"iss":  "init",
	})

	tokenString, err := t.SignedString([]byte(secret))

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

	t := c.Request().Header.Get("x-token")
	u, err := model.GetUser(h.db, t)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("ユーザ情報を取得しました")

	return c.JSON(http.StatusOK, UserGetResponse{Name: u.Name})
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

// ユーザ情報更新
func (h *Handler) UpdateUser(c echo.Context) error {

	req := new(UserUpdateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	n := req.Name
	t := c.Request().Header.Get("x-token")
	err := model.UpdateUser(h.db, n, t)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("ユーザ情報を更新しました")

	return c.NoContent(http.StatusOK)
}

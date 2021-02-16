package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type UserCreateRequest struct {
	Name string `json:"name"`
	// Token string `json:"token"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

// ユーザアカウント認証情報作成
func CreateUser(c echo.Context) error {

	req := new(UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	name := req.Name
	token := CreateToken(name)
	_, err := puser.Create(name, token)

	if err != nil {
		fmt.Println(err)
		return err
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

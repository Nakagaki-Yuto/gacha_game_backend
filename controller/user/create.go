package user

import (
	"fmt"
	"net/http"
	"log"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"

	"go_practice_mvc/model/user"
)

type User struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}


// ユーザアカウント認証情報作成
func CreateUser(c echo.Context) error {	

	u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
	name := u.Name
	token := CreateToken(name)
	error := puser.Create(name, token)

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
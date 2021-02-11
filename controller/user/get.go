package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"go_practice_mvc/model/user"
)

type UserGetResponse struct {
	Name string `json:"name"` 
}

// ユーザ情報取得
func GetUser(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user := User{}
	error := puser.Get(token, user)

	if error != nil {
		fmt.Println(error)
		return error
	} else {
		fmt.Println("ユーザ情報を取得しました")
	}

	return c.JSON(http.StatusOK, UserGetResponse{Name: user.Name})
}
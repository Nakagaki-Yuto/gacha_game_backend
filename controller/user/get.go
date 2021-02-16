package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"go_practice_mvc/model/user"
)

type UserGetResponse struct {
	Name string `json:"name"` 
}

// ユーザ情報取得
func GetUser(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user, error := puser.Get(token)

	if error != nil {
		fmt.Println(error)
		return error
	}

	fmt.Println("ユーザ情報を取得しました")
	
	return c.JSON(http.StatusOK, UserGetResponse{Name: user.Name})
}
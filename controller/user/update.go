package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	puser "go_practice_mvc/model/user"
)

// ユーザ情報更新
func UpdateUser(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	name := u.Name
	token := c.Request().Header.Get("x-token")
	error := puser.Update(name, token)
	fmt.Println(name)
	fmt.Println(token)
	fmt.Println(error)
	if error != nil {
		fmt.Println(error)
		return error
	}
	
	fmt.Println("ユーザ情報を更新しました")

	return c.JSON(http.StatusOK, "name change succeed")
}

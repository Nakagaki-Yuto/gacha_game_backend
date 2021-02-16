package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type UserGetResponse struct {
	Name string `json:"name"`
}

type Handler struct {
	db *DB
}

func New(db *DB) *Handler {
	return &Handler{
		db: db,
	}
}

// ユーザ情報取得
func (h *Handler) GetUser(c echo.Context) error {

	token := c.Request().Header.Get("x-token")
	user, _ := h.db.GetUser()
	// user, _ := puser.Get(token)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	fmt.Println("ユーザ情報を取得しました")

	return c.JSON(http.StatusOK, UserGetResponse{Name: user.Name})
}

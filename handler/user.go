package handler

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator"

	"go_practice_mvc/model"
)

type UserCreateRequest struct {
	Name string `validate:"required"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

// ユーザアカウント認証情報作成
func (h *Handler) CreateUser(c echo.Context) error {

	logger, _ := zap.NewDevelopment()
	validate := validator.New()

	req := new(UserCreateRequest)

	if err := c.Bind(req); err != nil {
		logger.Error("request is bad")
		return c.JSON(http.StatusBadRequest, MyError{Msg: err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		logger.Error("request is required")
		return c.JSON(http.StatusBadRequest, MyError{Msg: err.Error()})
	}

	n := req.Name

	if n==" " {
		logger.Error("empty string can't register db")
		return c.JSON(http.StatusBadRequest, MyError{Msg: "empty string can't register for db"})
	}
	
	t := CreateToken(n)
	err := model.CreateUser(h.db, n, t)

	if err != nil {
		fmt.Println("エラー3")
		return ErrorHandler(&err, c)
	}

	logger.Info("ユーザアカウント認証情報を作成しました")
	return c.JSON(http.StatusOK, UserCreateResponse{Token: t})
}

// ト-クンの作成
func CreateToken(name string) string {

	// 鍵となる文字列
	secret := "secret"

	// Token を作成
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, _ := t.SignedString([]byte(secret))

	return tokenString
}

type UserGetResponse struct {
	Name string `json:"name"`
}

// ユーザ情報取得
func (h *Handler) GetUser(c echo.Context) error {

	logger, _ := zap.NewDevelopment()
	t := c.Request().Header.Get("x-token")
	u, err := model.GetUser(h.db, t)

	if err != nil {
		return ErrorHandler(&err, c)
	}

	logger.Info("ユーザ情報を取得しました")
	return c.JSON(http.StatusOK, UserGetResponse{Name: u.Name})
}

type UserUpdateRequest struct {
	Name string `validate:"required"`
}

// ユーザ情報更新
func (h *Handler) UpdateUser(c echo.Context) error {

	logger, _ := zap.NewDevelopment()
	validate := validator.New()

	req := new(UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		logger.Error("request is bad")
		return c.JSON(http.StatusBadRequest, MyError{Msg: err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		logger.Error("request is required")
		return c.JSON(http.StatusBadRequest, MyError{Msg: err.Error()})
	}
	
	n := req.Name
	
	if n==" " {
		logger.Error("empty string can't register db")
		return c.JSON(http.StatusBadRequest, MyError{Msg: "empty string can't register db"})
	}


	t := c.Request().Header.Get("x-token")
	err := model.UpdateUser(h.db, n, t)
	
	if err != nil {
		return ErrorHandler(&err, c)
	}
	
	logger.Info("ユーザ情報を更新しました")
	return c.NoContent(http.StatusOK)
}

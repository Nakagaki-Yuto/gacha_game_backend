package handler

import (
	"net/http"
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handler {

	return &Handler{
		db: db,
	}
}

type MyError struct {
	Msg string `json:"message"`
}

func ErrorHandler(err *error, c echo.Context) error {
	logger, _ := zap.NewDevelopment()
	switch *err {
	case gorm.ErrRecordNotFound:
		logger.Error("token not found")
		return c.JSON(http.StatusNotFound, MyError{Msg: "token not found"})
	default:
		logger.Error("Internal Server Error")
		return c.JSON(http.StatusInternalServerError, MyError{Msg: (*err).Error()})
	}
}


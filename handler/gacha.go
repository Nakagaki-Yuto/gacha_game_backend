package handler

import (
	"math/rand"
	"net/http"
	"time"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator"

	"go_practice_mvc/model"
)

type GachaDrawRequest struct {
	Times int `validate:"required,min=1"`
}

type GachaDrawResponse struct {
	Results GachaResults `json:"results"`
}

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type GachaResults []GachaResult

// ガチャ実行
func (h *Handler) DrawGacha(c echo.Context) error {

	logger, _ := zap.NewDevelopment()
	validate := validator.New()
	
	t := c.Request().Header.Get("x-token")
	u, err := model.GetUserID(h.db, t)

	if err != nil {
		return ErrorHandler(&err, c)
	}

	uI := u.ID
	req := new(GachaDrawRequest)

	if err := c.Bind(req); err != nil {
		logger.Error("request is bad")
		return c.JSON(http.StatusBadRequest, MyError{Msg: err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		logger.Error("request is required")
		return c.JSON(http.StatusBadRequest, MyError{Msg: err.Error()})
	}

	times := req.Times

	var gachaResults GachaResults

	for i := 0; i < times; i++ {
		characterID, err := h.Gacha()

		if err != nil {
			return ErrorHandler(&err, c)
		}

		err = model.CreateUserCharacter(h.db, uI, characterID)

		if err != nil {
			return ErrorHandler(&err, c)
		}

		chara, err := model.GetCharacter(h.db, characterID)

		if err != nil {
			return ErrorHandler(&err, c)
		}

		gr := GachaResult{
			CharacterID: chara.ID,
			Name:        chara.Name,
		}

		gachaResults = append(gachaResults, gr)
	}

	logger.Info("ガチャを引きました")
	return c.JSON(http.StatusOK, GachaDrawResponse{Results: gachaResults})
}

// ガチャを引く
func (h Handler) Gacha() (string, error) {

	gachaRates, err := model.GetGachaRate(h.db)

	if err != nil {
		return "", err
	}

	kind := len(gachaRates)
	rates := make([]int, kind)
	maxCnt := 0

	for i := 0; i < kind; i++ {
		targetRate := gachaRates[i].Rate
		maxCnt += targetRate
		rates[i] = targetRate
	}

	parameter := make([]string, maxCnt)
	targetCnt := 0
	for i := 0; i < kind; i++ {
		for j := targetCnt; j < targetCnt+rates[i]; j++ {
			parameter[j] = gachaRates[i].CharacterID
		}
		targetCnt += rates[i]
	}

	rand.Seed(time.Now().UnixNano())

	return parameter[rand.Intn(maxCnt)], nil

}

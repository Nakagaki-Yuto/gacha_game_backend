package server
 
import (
	"github.com/labstack/echo"
	"go_practice_mvc/controller/character"
	"go_practice_mvc/controller/gacha"
	"go_practice_mvc/controller/user"
)

func NewRouter() *echo.Echo {
	
	e := echo.New()

	// ルーティング
	e.POST("/user/create", user.CreateUser)
	e.GET("/user/get", user.GetUser)
	e.PUT("/user/update", user.UpdateUser)
	e.POST("/gacha/draw", gacha.DrawGacha)
	e.GET("/character/list", character.GetCharacterList)
	
    return e
}


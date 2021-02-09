package server
 
import (
	"github.com/labstack/echo"
	"go_practice_mvc/handler/character"
	"go_practice_mvc/handler/gacha"
	"go_practice_mvc/handler/user"
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


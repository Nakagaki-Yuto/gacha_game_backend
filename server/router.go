package server
 
import (
	"github.com/labstack/echo/v4/middleware" 
	"github.com/labstack/echo/v4"

	"go_practice_mvc/controller/character"
	"go_practice_mvc/controller/gacha"
	"go_practice_mvc/controller/user"
)

func NewRouter() *echo.Echo {
	
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:3000/", "http://localhost:8080/"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//   }))

	// ルーティング
	e.POST("/user/create", user.CreateUser)
	e.GET("/user/get", user.GetUser)
	e.PUT("/user/update", user.UpdateUser)
	e.POST("/gacha/draw", gacha.DrawGacha)
	e.GET("/character/list", character.GetCharacterList)
	
    return e
}


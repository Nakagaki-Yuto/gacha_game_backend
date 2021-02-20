package server
 
import (
	"github.com/labstack/echo/v4/middleware" 
	"github.com/labstack/echo/v4"

	"go_practice_mvc/handler"
)



type Server struct {
	handler handler.Handler
}

func New(handler *handler.Handler) *Server {
	return &Server{
		handler: *handler,
	}
}

func (s *Server) Run() error {
	
	e := echo.New()
	e.Use(middleware.CORS())

	// ルーティング
	e.POST("/user/create", s.handler.CreateUser)
	e.GET("/user/get", s.handler.GetUser)
	e.PUT("/user/update", s.handler.UpdateUser)
	e.POST("/gacha/draw", s.handler.DrawGacha)
	e.GET("/character/list", s.handler.GetCharacterList)
	e.GET("/ranking", s.handler.GetRanking)

	return e.Start(":8080")
}


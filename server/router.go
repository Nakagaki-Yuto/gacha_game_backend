package server

import (
	"github.com/labstack/echo"
)

type Server struct {
	handler Handler
}

func New(db *DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Run() error {
	e := echo.New()

	// ルーティング
	e.POST("/user/create", s.handler.CreateUser)
	e.GET("/user/get", s.handler.GetUser)
	e.PUT("/user/update", s.handler.UpdateUser)
	e.POST("/gacha/draw", s.handler.DrawGacha)
	e.GET("/character/list", s.handler.GetCharacterList)

	return e.Start(":8080")
}

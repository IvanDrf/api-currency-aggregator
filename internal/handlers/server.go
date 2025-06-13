package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Server *echo.Echo
}

func InitServer() Server {
	return Server{
		Server: echo.New(),
	}
}

func (s *Server) Start() {
	if err := s.Server.Start(":8080"); err != nil {
		log.Fatal("can't start server")
	}
}

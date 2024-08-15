package server

import (
	"consume-api-soap-rest/internal/delivery/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {

	group := s.engine.Group("/api/v1")

	handler.NewConsumeCalculator(group).Route()
	handler.NewConsumeRestEvent(group).Route()

}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {

	engine := gin.Default()
	return &Server{engine: engine, host: ":8090"}
}

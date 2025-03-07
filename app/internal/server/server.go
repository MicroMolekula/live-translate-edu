package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/controllers"
	"log"
)

type Server struct {
	engine *gin.Engine
	cfg    *configs.Config
	router *controllers.Router
}

func NewServer(cfg *configs.Config, router *controllers.Router) *Server {
	engine := gin.Default()
	return &Server{
		engine: engine,
		cfg:    cfg,
		router: router,
	}
}

func (s *Server) Run() {
	s.router.InitRoutes(s.engine)
	if err := s.engine.Run(fmt.Sprintf(":%d", s.cfg.ServerPort)); err != nil {
		log.Fatal(err)
	}
}

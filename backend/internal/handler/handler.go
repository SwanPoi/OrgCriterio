package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SwanPoi/OrgCriterio/internal/logger"
	"github.com/SwanPoi/OrgCriterio/internal/usecase"
)

type Server struct {
	router		*gin.Engine
	log			logger.Logger
	uc			usecase.Usecase
}

func NewServer(log logger.Logger, uc usecase.Usecase) *Server {
	router := gin.New()

	router.Use(LogMiddleware(log), gin.Recovery())

	server := &Server{
		router: router,
		log: log,
		uc: uc,
	}

	server.RegisterRoutes()

	return server
}

func (s *Server) RegisterRoutes() {
	s.router.GET("/health", func(c *gin.Context) {
		s.log.Debug("health check requested")
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}
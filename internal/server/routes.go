package server

import (
	"github.com/bismastr/discord-bot/internal/handler/rest"
	"github.com/bismastr/discord-bot/internal/session"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() {
	apiV1 := s.router.Group("api/v1")

	s.healthRoutes(apiV1)
	s.gamingSessionRoutes(apiV1)
	s.botGamingSessionRoutes(apiV1)
}

func (s *Server) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")

	{
		h := rest.NewHealthCtrl()
		healthRoutes.GET("/ping", h.Ping)
	}
}

func (s *Server) gamingSessionRoutes(api *gin.RouterGroup) {
	gamingSessionRoutes := api.Group("/gaming-session")

	{
		repository := session.NewRepositoryImpl(s.database)
		h := rest.NewSessionCtrl(session.NewGamingSessionService(repository))

		gamingSessionRoutes.POST("/", h.CreateGamingSession)
		gamingSessionRoutes.PUT("/", h.UpdateGamingSessionByRefId)
	}
}

func (s *Server) botGamingSessionRoutes(api *gin.RouterGroup) {
	botRoutes := api.Group("/bot/gaming-session")

	{
		repository := session.NewRepositoryImpl(s.database)
		h := rest.NewBotCtrl(session.NewBotGamingSessionService(repository, s.dg), session.NewGamingSessionService(repository))

		botRoutes.POST("/create", h.CreateGamingSession)
	}
}

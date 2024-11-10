package server

import (
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/internal/handler"
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
		h := handler.NewHealthCtrl()
		healthRoutes.GET("/ping", h.Ping)
	}
}

func (s *Server) gamingSessionRoutes(api *gin.RouterGroup) {
	gamingSessionRoutes := api.Group("/gaming-session")

	{
		repository := gamingSession.NewRepositoryImpl(s.database)
		h := handler.NewSessionCtrl(gamingSession.NewGamingSessionService(repository))

		gamingSessionRoutes.POST("/", h.CreateGamingSession)
		gamingSessionRoutes.PUT("/", h.UpdateGamingSessionByRefId)
	}
}

func (s *Server) botGamingSessionRoutes(api *gin.RouterGroup) {
	botRoutes := api.Group("/bot/gaming-session")

	{
		repository := gamingSession.NewRepositoryImpl(s.database)
		h := handler.NewBotCtrl(bot.NewBotGamingSessionService(repository, s.dg), gamingSession.NewGamingSessionService(repository))

		botRoutes.POST("/create", h.CreateGamingSession)
	}
}

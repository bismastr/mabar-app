package server

import (
	"github.com/bismastr/discord-bot/internal/handler"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes(h *handler.Handler) {
	apiV1 := s.router.Group("api/v1")

	s.gamingSessionRoutes(apiV1, h)
	s.botGamingSessionRoutes(apiV1, h)
	s.authRoutes(apiV1, h)
}

func (s *Server) botGamingSessionRoutes(api *gin.RouterGroup, h *handler.Handler) {
	botRoutes := api.Group("/bot/gaming-session")

	{
		botRoutes.POST("/create", h.CreateGamingSession)
	}
}
func (s *Server) gamingSessionRoutes(api *gin.RouterGroup, h *handler.Handler) {
	gamingSessionRoutes := api.Group("/gaming-session")

	{
		gamingSessionRoutes.POST("/create", h.CreateGamingSessionV2)
		gamingSessionRoutes.POST("/join", h.JoinGamingSession)
		gamingSessionRoutes.GET("/:id", h.GetGamingSession)
		gamingSessionRoutes.GET("/", h.GetAllGamingSessions)
	}
}

func (s *Server) authRoutes(api *gin.RouterGroup, h *handler.Handler) {
	authRoutes := api.Group("/auth")

	{
		authRoutes.GET("/:provider/callback", h.Callback)
		authRoutes.GET("/:provider", h.Login)
		authRoutes.GET("/profile", h.CheckIsAuthenticaed)
	}
}

package server

import (
	"time"

	"github.com/bismastr/discord-bot/internal/config"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	dg     *discordgo.Session
}

func NewServer(e *gin.Engine, dg *discordgo.Session) *Server {
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:5173", "https://mabar.bism.app", "https://app-mabar.bism.app"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return &Server{
		router: e,
		dg:     dg,
	}
}

func (s *Server) Start() error {
	return s.router.Run(config.Envs.Port)
}

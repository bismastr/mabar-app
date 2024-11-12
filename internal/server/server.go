package server

import (
	"github.com/bismastr/discord-bot/internal/config"
	"github.com/bismastr/discord-bot/internal/database"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	database *database.DbClient
	dg       *discordgo.Session
}

func NewServer(e *gin.Engine, db *database.DbClient, dg *discordgo.Session) *Server {
	return &Server{
		router:   e,
		database: db,
		dg:       dg,
	}
}

func (s *Server) Start() error {
	return s.router.Run(config.Envs.Port)
}

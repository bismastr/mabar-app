package server

import (
	"github.com/bismastr/discord-bot/internal/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	database *database.DbClient
}

func NewServer(e *gin.Engine, db *database.DbClient) *Server {
	return &Server{
		router:   e,
		database: db,
	}

}

func (s *Server) Start() error {
	return s.router.Run(":8080")
}

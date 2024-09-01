package server

import (
	"net/http"
	"time"

	"github.com/bismastr/discord-bot/db"
)


type Server struct {
	db db.DbClient
}

func NewServer(database *db.DbClient) *http.Server {
	NewServer := &Server{
		db: *database,
	}

	server := &http.Server{
		Addr: ":8080",
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
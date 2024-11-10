package handler

import (
	"net/http"

	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/gin-gonic/gin"
)

type sessionCtrl struct {
	sessionService *gamingSession.GamingSessionService
}

func NewSessionCtrl(sessionService *gamingSession.GamingSessionService) *sessionCtrl {
	return &sessionCtrl{
		sessionService: sessionService,
	}
}

func (s *sessionCtrl) CreateGamingSession(c *gin.Context) {

	newGamingSession := gamingSession.GamingSession{
		GameName: "mabar",
		CreatedBy: &gamingSession.CreatedBy{
			Id:       "1",
			Username: "bismastr",
		},
	}

	s.sessionService.CreateGamingSession(c.Request.Context(), newGamingSession)

	c.JSON(http.StatusOK, gin.H{"message": "gaming session created"})
}

func (s *sessionCtrl) UpdateGamingSessionByRefId(c *gin.Context) {

	newGamingSession := gamingSession.GamingSession{
		GameName: "mabar 2 update testing2",
		CreatedBy: &gamingSession.CreatedBy{
			Id:       "1",
			Username: "testupdate 1",
		},
	}

	s.sessionService.UpdateGamingSessionByRefId(c.Request.Context(), "A0BZAFQYYpYKO4utwVHQ", newGamingSession)

	c.JSON(http.StatusOK, gin.H{"message": "gaming session update"})
}

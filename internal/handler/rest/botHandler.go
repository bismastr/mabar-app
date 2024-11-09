package rest

import (
	"net/http"

	"github.com/bismastr/discord-bot/internal/session"
	"github.com/gin-gonic/gin"
)

type botCtrl struct {
	botService     *session.BotGamingSessionService
	sessionService *session.GamingSessionService
}

func NewBotCtrl(b *session.BotGamingSessionService, s *session.GamingSessionService) *botCtrl {
	return &botCtrl{
		botService:     b,
		sessionService: s,
	}
}

func (b *botCtrl) CreateGamingSession(c *gin.Context) {
	var newGamingSession session.GamingSession

	if err := c.BindJSON(&newGamingSession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	id, err := b.sessionService.CreateGamingSession(c, newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	res, err := b.botService.CreateGamingSession(id, newGamingSession.GameName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

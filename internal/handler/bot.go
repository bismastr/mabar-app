package handler

import (
	"net/http"

	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/gin-gonic/gin"
)

type botCtrl struct {
	botService     *bot.BotGamingSessionService
	sessionService *gamingSession.GamingSessionService
}

func NewBotCtrl(b *bot.BotGamingSessionService, s *gamingSession.GamingSessionService) *botCtrl {
	return &botCtrl{
		botService:     b,
		sessionService: s,
	}
}

func (b *botCtrl) CreateGamingSession(c *gin.Context) {
	var newGamingSession gamingSession.GamingSession

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

	res, err := b.botService.CreateGamingSession(id, &newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

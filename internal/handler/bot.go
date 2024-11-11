package handler

import (
	"net/http"

	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGamingSession(c *gin.Context) {
	var newGamingSession gamingSession.GamingSession

	if err := c.BindJSON(&newGamingSession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	id, err := h.gamingSession.CreateGamingSession(c, newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	res, err := h.bot.CreateGamingSession(id, &newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

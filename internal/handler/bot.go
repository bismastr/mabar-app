package handler

import (
	"fmt"
	"net/http"

	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGamingSession(c *gin.Context) {
	var newGamingSession gaming_session.CreateGamingSessionRequest
	if err := c.BindJSON(&newGamingSession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result, err := h.gaming_session.CreateGamingSession(c, &newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	response, err := h.gaming_session.GetGamingSessionById(c, result.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(newGamingSession.ChannelID)

	res, err := h.bot.CreateGamingSession(response, newGamingSession.ChannelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

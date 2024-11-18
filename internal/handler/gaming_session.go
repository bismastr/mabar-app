package handler

import (
	"net/http"
	"time"

	"github.com/bismastr/discord-bot/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSession(c *gin.Context) {
	newGamingSession := &model.Session{
		IsFinish:  false,
		CreatedAt: time.Now(),
		CreatedBy: 364635434304798730,
		GameId:    1,
	}

	res, err := h.gaming.CreateGaming(newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

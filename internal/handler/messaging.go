package handler

import (
	"github.com/bismastr/discord-bot/internal/notification"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendMessageWithTopic(c *gin.Context) {
	var request *notification.SendNotificationWithTopicRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.notification.SendMessage(c, request); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success send message"})
}

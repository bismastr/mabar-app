package handler

import (
	"net/http"
	"strconv"

	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGamingSessionV2(c *gin.Context) {
	var newGamingSession *gaming_session.CreateGamingSessionRequest
	if err := c.BindJSON(&newGamingSession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := h.gaming_session.CreateGamingSession(c, newGamingSession)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) JoinGamingSession(c *gin.Context) {
	var joinGamingSesionRequest *gaming_session.JoinGamingSesionRequest
	if err := c.BindJSON(&joinGamingSesionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err := h.gaming_session.InsertUserJoinSession(c, joinGamingSesionRequest.UserId, joinGamingSesionRequest.SessionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "joined session",
	})
}

func (h *Handler) GetGamingSession(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	result, err := h.gaming_session.GetGamingSessionById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetAllGamingSessions(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	rows, err := strconv.Atoi(c.DefaultQuery("rows", "16"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	req := gaming_session.GetAllGamingSessionRequest{
		Page: page,
		Rows: rows,
	}

	result, err := h.gaming_session.GetAllGamingSessions(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetAllGames(c *gin.Context) {
	result, err := h.gaming_session.GetAllGames(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

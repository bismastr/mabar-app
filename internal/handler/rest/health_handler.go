package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCtrl struct{}

func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}

func (h *healthCtrl) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "service ok"})
}

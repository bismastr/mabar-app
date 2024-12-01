package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *Handler) Prometheus(c *gin.Context) {
	p := promhttp.Handler()

	p.ServeHTTP(c.Writer, c.Request)
}

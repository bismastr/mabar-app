package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler{
	r := gin.Default()

	r.GET("/health", s.HealthCheck)

	return r
}

func (s *Server) HealthCheck(c *gin.Context)  {
	resp := make(map[string]string)
	resp["message"] = "service ok"

	c.JSON(http.StatusOK, resp)
}
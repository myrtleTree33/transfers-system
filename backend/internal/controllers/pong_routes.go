package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongResponse struct {
	Message string `json:"message"`
}

// @Summary      Ping handler
// @Description  Ping handler
// @Produce      json
// @Success      200  {object}  PongResponse
// @Router       /v1/ping [get]
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, PongResponse{
		Message: "pong",
	})
}

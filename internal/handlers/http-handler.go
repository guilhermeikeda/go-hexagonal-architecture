package handlers

import (
	"com.guilherme/hexagonal-example/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	gameService ports.GameService
}

func NewHandler(gameService ports.GameService) *HTTPHandler {
	return &HTTPHandler{gameService}
}

func (httpHandler *HTTPHandler) Get(c *gin.Context) {
	game, err := httpHandler.gameService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, game)
}

func (httpHandler *HTTPHandler) Create(c *gin.Context) {
	// httpHandler.gameService.Create()
}

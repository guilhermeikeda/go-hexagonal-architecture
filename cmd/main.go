package main

import (
	"com.guilherme/hexagonal-example/internal/core/services"
	"com.guilherme/hexagonal-example/internal/handlers"
	repository "com.guilherme/hexagonal-example/internal/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := repository.NewInMemoryRepository()
	gamesService := services.New(gamesRepository)
	gamesHandler := handlers.NewHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)

	router.Run(":8080")
}

package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Iniciando api...")

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8008")
}

package main

import (
	"log"
	"log/slog"

	"github.com/Pedro-0101/peridot/internal/db"
	"github.com/Pedro-0101/peridot/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.Connect()
	if err != nil {
		slog.Error("Falha ao conectar no banco", "error", err)
		panic(err)
	}
	defer dbConnection.Close()

	routes.InitRoutes(&server.RouterGroup, dbConnection)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := server.Run(":8008"); err != nil {
		log.Fatal(err)
	}
}

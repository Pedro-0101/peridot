package main

import (
	"log/slog"

	"github.com/Pedro-0101/peridot/internal/controllers"
	"github.com/Pedro-0101/peridot/internal/db"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"github.com/Pedro-0101/peridot/internal/usecases"
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

	userRepo := repositories.NewUserRepository(dbConnection)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userCtrl := controllers.NewUserController(userUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/users", userCtrl.CreateUser)
	server.GET("/users", userCtrl.GetAllUsers)

	server.Run(":8008")
}

package routes

import (
	"database/sql"

	"github.com/Pedro-0101/peridot/internal/controllers"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"github.com/Pedro-0101/peridot/internal/usecases"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, dbConnection *sql.DB) {

	userRepo := repositories.NewUserRepository(dbConnection)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userCtrl := controllers.NewUserController(userUseCase)

	r.POST("/users", userCtrl.CreateUser)
	r.GET("/users", userCtrl.GetAllUsers)
	r.GET("/user:id", userCtrl.GetUserById)
}

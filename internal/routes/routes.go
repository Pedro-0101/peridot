package routes

import (
	"database/sql"

	"github.com/Pedro-0101/peridot/internal/controllers"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"github.com/Pedro-0101/peridot/internal/services/user_service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, dbConnection *sql.DB) {

	userRepo := repositories.NewUserRepository(dbConnection)
	userService := user_service.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userService)

	r.POST("/users", userCtrl.CreateUser)
	r.GET("/users", userCtrl.GetAllUsers)
	r.GET("/user:id", userCtrl.GetUserById)
	r.PUT("/user:id", userCtrl.UpdateUser)
	r.DELETE("/user:id", userCtrl.DeleteUser)
}

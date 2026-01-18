package controllers

import (
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserController(uuc *usecases.UserUseCase) *UserController {
	return &UserController{UserUseCase: uuc}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var u users.User

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if err := c.UserUseCase.CreateUser(&u); err != nil {
		ctx.JSON(500, gin.H{"error": "Error on create user"})
		return
	}

	ctx.JSON(201, u)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserUseCase.GetAllUsers()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error on search users"})
	}

	ctx.JSON(200, users)
}

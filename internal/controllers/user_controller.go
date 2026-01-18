package controllers

import (
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	CreateUserUC  *usecases.CreateUserUseCase
	GetAllUsersUC *usecases.GetAllUsersUseCase
}

func NewUserController(createUsercaUC *usecases.CreateUserUseCase, getAllUsers *usecases.GetAllUsersUseCase) *UserController {
	return &UserController{CreateUserUC: createUsercaUC, GetAllUsersUC: getAllUsers}
}

func (c *UserController) Create(ctx *gin.Context) {
	var u users.User

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if err := c.CreateUserUC.Execute(&u); err != nil {
		ctx.JSON(500, gin.H{"error": "Error on create user"})
		return
	}

	ctx.JSON(201, u)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.GetAllUsersUC.Execute()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error on search users"})
	}

	ctx.JSON(200, users)
}

package controllers

import (
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func NewUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (c *UserController) Create(ctx *gin.Context) {
	var u users.User

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := c.Repo.CreateUser(&u); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao salvar usuário"})
		return
	}

	ctx.JSON(201, u)
}

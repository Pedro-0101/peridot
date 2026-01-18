package controllers

import (
	"fmt"
	"net/http"

	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/request"
	"github.com/Pedro-0101/peridot/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserController(uuc *usecases.UserUseCase) *UserController {
	return &UserController{UserUseCase: uuc}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var u request.UserRequest

	if err := ctx.ShouldBindJSON(&u); err != nil {
		resterr := resterr.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields: %s", err))
		ctx.JSON(resterr.Code, resterr)
		return
	}

	responseUser, err := c.UserUseCase.CreateUser(&u)

	if err != nil {
		resterr := resterr.NewInternalServerError(
			fmt.Sprintf("An internal server error has ocurred: %s", err))
		ctx.JSON(resterr.Code, resterr)
		return
	}

	ctx.JSON(201, responseUser)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserUseCase.GetAllUsers()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error on search users"})
	}

	ctx.JSON(200, users)
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.UserUseCase.GetUserById(id)

	if err != nil {
		if err.Error() == "Error: Error searching for user" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if user.ID == uuid.Nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

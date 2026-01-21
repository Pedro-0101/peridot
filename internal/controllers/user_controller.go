package controllers

import (
	"net/http"

	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/request"
	"github.com/Pedro-0101/peridot/internal/services/user_service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService user_service.UserDomainService
}

func NewUserController(us user_service.UserDomainService) *UserController {
	return &UserController{UserService: us}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var u request.UserRequest

	if err := ctx.ShouldBindJSON(&u); err != nil {
		errRest := resterr.NewBadRequestError("Invalid JSON fields")
		ctx.JSON(errRest.Code, errRest)
		return
	}

	responseUser, err := c.UserService.CreateUser(u)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, responseUser)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserService.GetAllUsers()

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(200, users)
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.UserService.GetUserById(id)
	if err != nil {

		if err.Code == 404 {
			ctx.JSON(err.Code, "User not found")
			return
		}

		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var u request.UserRequest

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.UserService.GetUserById(id)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	newUser, err := c.UserService.UpdateUser(id, u)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.UserService.DeleteUser(id)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, "User deleted successfully")
}

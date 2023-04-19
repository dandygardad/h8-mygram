package controllers

import (
	"github.com/gin-gonic/gin"
	"mygram/helper"
	"mygram/model/entity"
	"mygram/model/web"
	"net/http"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

// Register godoc
// @Summary Register User
// @Description Create an account for user
// @tags user
// @Accept json
// @Produce json
// @Param user body web.RegisterRequest true "Create an account for user"
// @Success 201 {object} entity.User
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /user/register [post]
func (c *Controller) Register(ctx *gin.Context) {
	var request web.RegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helper.ResponseError(ctx, "invalid json", http.StatusBadRequest)
		return
	}

	// Validation
	err = request.Validation()
	if err != nil {
		helper.ValidationError(ctx, err)
		return
	}

	registerUser := entity.User{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Age:      request.Age,
	}

	created, err := c.user.Create(registerUser)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

// Login godoc
// @Summary Login User
// @Description Login for user
// @tags user
// @Accept json
// @Produce json
// @Param user body web.LoginRequest true "Login for user"
// @Success 200 {object} web.ResponseLogin
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /user/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	var request web.LoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helper.ResponseError(ctx, "invalid json", http.StatusBadRequest)
		return
	}

	// Validation
	err = request.Validation()
	if err != nil {
		helper.ValidationError(ctx, err)
		return
	}

	login := entity.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := c.user.Login(login)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	token := helper.GenerateToken(user)

	ctx.JSON(http.StatusOK, web.ResponseLogin{
		Token: token,
		Data:  user,
	})
}

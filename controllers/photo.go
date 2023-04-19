package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mygram/helper"
	"mygram/model/entity"
	"mygram/model/web"
	"net/http"
	"strconv"
)

type PhotoController interface {
	GetAllPhoto(ctx *gin.Context)
	GetOnePhoto(ctx *gin.Context)
	CreatePhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

// GetAllPhoto godoc
// @Summary Get All Photo
// @Description Get every photo
// @tags photo
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /photo [get]
func (c *Controller) GetAllPhoto(ctx *gin.Context) {
	results, err := c.photo.GetAllPhoto()
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

// GetOnePhoto godoc
// @Summary Get One Photo
// @Description Get photo by id
// @tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID photo"
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /photo/{id} [get]
func (c *Controller) GetOnePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	find := entity.Photo{
		ID: uint(cvtId),
	}

	result, err := c.photo.GetOnePhoto(find)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// CreatePhoto godoc
// @Summary Create Photo
// @Description Create new photo
// @tags photo
// @Accept json
// @Produce json
// @Param photo body web.PhotoRequest true "Request for photo"
// @Param authorization header string true "Token" default(Bearer <insert-token>)
// @Success 201 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /photo/create [post]
func (c *Controller) CreatePhoto(ctx *gin.Context) {
	var request web.PhotoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helper.ResponseError(ctx, "invalid json", http.StatusBadRequest)
		return
	}

	// Validation
	err = request.Validation()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		helper.ValidationError(ctx, err)
		return
	}

	// Get data from jwt
	jwtData, _ := ctx.Get("userData")
	cvtJwtData := jwtData.(jwt.MapClaims)

	create := entity.Photo{
		Title:    request.Title,
		PhotoURL: request.PhotoURL,
		Caption:  request.Caption,
		UserID:   uint(cvtJwtData["id"].(float64)),
	}

	result, err := c.photo.CreatePhoto(create)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// UpdatePhoto godoc
// @Summary Update Photo
// @Description Update photo by id
// @tags photo
// @Accept json
// @Produce json
// @Param photo body web.PhotoRequest true "Request for photo"
// @Param authorization header string true "Token" default(Bearer <insert-token>)
// @Param id path int true "ID photo"
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /photo/update/{id} [put]
func (c *Controller) UpdatePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	var request web.PhotoRequest
	err = ctx.ShouldBindJSON(&request)
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

	// Get data from jwt
	jwtData, _ := ctx.Get("userData")
	cvtJwtData := jwtData.(jwt.MapClaims)

	update := entity.Photo{
		Title:    request.Title,
		PhotoURL: request.PhotoURL,
		Caption:  request.Caption,
		UserID:   uint(cvtJwtData["id"].(float64)),
	}

	result, err := c.photo.UpdatePhoto(update, cvtId)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// DeletePhoto godoc
// @Summary Delete Photo
// @Description Delete photo by id
// @tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID photo"
// @Param authorization header string true "Token" default(Bearer <insert-token>)
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /photo/delete/{id} [delete]
func (c *Controller) DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	// Get data from jwt
	jwtData, _ := ctx.Get("userData")
	cvtJwtData := jwtData.(jwt.MapClaims)

	inputPhoto := entity.Photo{
		UserID: uint(cvtJwtData["id"].(float64)),
	}

	err = c.photo.DeletePhoto(inputPhoto, cvtId)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}

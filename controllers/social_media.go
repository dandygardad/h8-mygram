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

type SocialMediaController interface {
	GetAllSocialMedia(ctx *gin.Context)
	GetOneSocialMedia(ctx *gin.Context)
	CreateSocialMedia(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

func (c *Controller) GetAllSocialMedia(ctx *gin.Context) {
	results, err := c.socmed.GetAllSocialMedia()
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *Controller) GetOneSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	find := entity.SocialMedia{
		ID: uint(cvtId),
	}

	result, err := c.socmed.GetOneSocialMedia(find)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Controller) CreateSocialMedia(ctx *gin.Context) {
	var request web.SocialMediaRequest
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

	create := entity.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         uint(cvtJwtData["id"].(float64)),
	}

	result, err := c.socmed.CreateSocialMedia(create)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *Controller) UpdateSocialMedia(ctx *gin.Context) {
	var request web.SocialMediaRequest
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

	// Get data from jwt
	jwtData, _ := ctx.Get("userData")
	cvtJwtData := jwtData.(jwt.MapClaims)

	update := entity.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         uint(cvtJwtData["id"].(float64)),
	}

	result, err := c.socmed.UpdateSocialMedia(update)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Controller) DeleteSocialMedia(ctx *gin.Context) {
	// Get data from jwt
	jwtData, _ := ctx.Get("userData")
	cvtJwtData := jwtData.(jwt.MapClaims)

	err := c.socmed.DeleteSocialMedia(int(cvtJwtData["id"].(float64)))
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}

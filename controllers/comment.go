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

type CommentController interface {
	GetAllComment(ctx *gin.Context)
	GetOneComment(ctx *gin.Context)
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

// GetAllComment godoc
// @Summary Get All Comment
// @Description Get every comment from photo id
// @tags comment
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /comment/all/{id} [get]
func (c *Controller) GetAllComment(ctx *gin.Context) {
	// Get all comments from photo id
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "photo id not valid", http.StatusBadRequest)
		return
	}

	results, err := c.comment.GetAllComment(cvtId)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

// GetOneComment godoc
// @Summary Get One Comment
// @Description Get comment from id
// @tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /comment/{id} [get]
func (c *Controller) GetOneComment(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID comment not valid", http.StatusBadRequest)
		return
	}

	find := entity.Comment{
		Id: uint(cvtId),
	}

	result, err := c.comment.GetOneComment(find)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// CreateComment godoc
// @Summary Create Comment
// @Description Create comment from photo id
// @tags comment
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Param authorization header string true "Token" default(Bearer <insert-token>)
// @Param comment body web.CommentRequest true "Request for comment"
// @Success 201 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /comment/create/{id} [post]
func (c *Controller) CreateComment(ctx *gin.Context) {
	// Ambil photo id
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID comment not valid", http.StatusBadRequest)
		return
	}

	// Ambil message
	var request web.CommentRequest
	err = ctx.ShouldBindJSON(&request)
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

	create := entity.Comment{
		Message: request.Message,
		PhotoID: uint(cvtId),
		UserID:  uint(cvtJwtData["id"].(float64)),
	}

	result, err := c.comment.CreateComment(create)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// UpdateComment godoc
// @Summary Update Comment
// @Description Update comment from id
// @tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param authorization header string true "Token" default(Bearer <insert-token>)
// @Param comment body web.CommentRequest true "Request for comment"
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /comment/update/{id} [put]
func (c *Controller) UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("id")
	cvtCommentId, err := strconv.Atoi(commentId)
	if err != nil {
		helper.ResponseError(ctx, "comment id not valid", http.StatusBadRequest)
		return
	}

	var request web.CommentRequest
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

	update := entity.Comment{
		UserID:  uint(cvtJwtData["id"].(float64)),
		Message: request.Message,
	}

	result, err := c.comment.UpdateComment(update, cvtCommentId)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// DeleteComment godoc
// @Summary Delete Comment
// @Description Delete comment from id
// @tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param authorization header string true "Token" default(Bearer <insert-token>)
// @Success 200 {object} entity.Photo
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /comment/delete/{id} [delete]
func (c *Controller) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	// Get data from jwt
	jwtData, _ := ctx.Get("userData")
	cvtJwtData := jwtData.(jwt.MapClaims)

	inputComment := entity.Comment{
		UserID: uint(cvtJwtData["id"].(float64)),
	}

	err = c.comment.DeleteComment(inputComment, cvtId)
	if err != nil {
		helper.CustomErrorMsg(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}

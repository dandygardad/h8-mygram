package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

func ResponseError(ctx *gin.Context, msg string, code int) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"message": msg,
	})
}

func CustomErrorMsg(ctx *gin.Context, err error) {
	switch err.Error() {
	case "username_already_exist":
		ResponseError(ctx, "Username already exist", http.StatusBadRequest)
	case "not_found":
		ResponseError(ctx, "ID not found", http.StatusBadRequest)
	case "no_data":
		ResponseError(ctx, "no data", http.StatusOK)
	case "photo_not_found":
		ResponseError(ctx, "photo not found", http.StatusBadRequest)
	case "already_exist":
		ResponseError(ctx, "already exist", http.StatusBadRequest)
	case "email_already_exist":
		ResponseError(ctx, "Email already exist", http.StatusBadRequest)
	case "user_not_exists":
		ResponseError(ctx, "Username not registered", http.StatusBadRequest)
	case "wrong_password":
		ResponseError(ctx, "Wrong password", http.StatusBadRequest)
	default:
		fmt.Println("Error:", err.Error())
		ResponseError(ctx, "Server Error", http.StatusInternalServerError)
	}
}

func ValidationError(ctx *gin.Context, err error) {
	var errorSlice []gin.H
	if e, ok := err.(validation.Errors); ok {
		for field, msg := range e {
			errorSlice = append(errorSlice, gin.H{
				field: msg.Error(),
			})
		}
	}
	ctx.AbortWithStatusJSON(400, gin.H{
		"validation": errorSlice,
	})
}

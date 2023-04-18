package routes

import (
	"github.com/gin-gonic/gin"
	"mygram/controllers"
	"mygram/services"
)

func UserRoutes(router *gin.Engine, service services.UserInterface) {
	handler := controllers.NewUserController(service)
	api := router.Group("/user")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)
	}
}

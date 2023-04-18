package routes

import (
	"github.com/gin-gonic/gin"
	"mygram/controllers"
	"mygram/middlewares"
	"mygram/services"
)

func SocialMediaRoutes(router *gin.Engine, service services.SocialMediaInterface) {
	handler := controllers.NewSocialMediaController(service)

	api := router.Group("/socmed")
	{
		api.GET("", handler.GetAllSocialMedia)
		api.GET("/:id", handler.GetOneSocialMedia)
		api.Use(middlewares.Authentication())
		api.POST("/create", handler.CreateSocialMedia)
		api.PUT("/update", handler.UpdateSocialMedia)
		api.DELETE("/delete", handler.DeleteSocialMedia)
	}
}

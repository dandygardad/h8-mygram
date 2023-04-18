package routes

import (
	"github.com/gin-gonic/gin"
	"mygram/controllers"
	"mygram/middlewares"
	"mygram/services"
)

func PhotoRoutes(router *gin.Engine, services services.PhotoInterface) {
	handler := controllers.NewPhotoController(services)

	api := router.Group("photo")
	{
		api.GET("", handler.GetAllPhoto)
		api.GET("/:id", handler.GetOnePhoto)
		api.Use(middlewares.Authentication())
		api.POST("/create", handler.CreatePhoto)
		api.PUT("/update/:id", handler.UpdatePhoto)
		api.DELETE("/delete/:id", handler.DeletePhoto)
	}
}

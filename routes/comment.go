package routes

import (
	"github.com/gin-gonic/gin"
	"mygram/controllers"
	"mygram/middlewares"
	"mygram/services"
)

func CommentRoutes(router *gin.Engine, services services.CommentInterface) {
	handler := controllers.NewCommentController(services)

	api := router.Group("comment")
	{
		api.GET("all/:id", handler.GetAllComment)
		api.GET("/:id", handler.GetOneComment)
		api.Use(middlewares.Authentication())
		api.POST("/create/:id", handler.CreateComment)
		api.PUT("/update/:id", handler.UpdateComment)
		api.DELETE("/delete/:id", handler.DeleteComment)
	}
}

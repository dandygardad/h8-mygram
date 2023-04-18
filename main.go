package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"mygram/config"
	"mygram/repositories"
	"mygram/routes"
	"mygram/services"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	// Initialize gorm and postgres
	err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	repo := repositories.NewUserRepo(config.NewGorm.DB)
	servUser := services.NewUserService(repo)
	servSocialMedia := services.NewSocialMediaService(repo)
	servPhoto := services.NewPhotoService(repo)
	servComment := services.NewCommentService(repo)

	newRouter := gin.New()
	routes.UserRoutes(newRouter, servUser)
	routes.SocialMediaRoutes(newRouter, servSocialMedia)
	routes.PhotoRoutes(newRouter, servPhoto)
	routes.CommentRoutes(newRouter, servComment)
	newRouter.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Route not found",
		})
	})

	port := os.Getenv("PORT")
	err = newRouter.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

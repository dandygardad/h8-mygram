package controllers

import "mygram/services"

type Controller struct {
	user    services.UserInterface
	socmed  services.SocialMediaInterface
	photo   services.PhotoInterface
	comment services.CommentInterface
}

func NewUserController(user services.UserInterface) *Controller {
	return &Controller{user: user}
}

func NewSocialMediaController(socmed services.SocialMediaService) *Controller {
	return &Controller{socmed: socmed}
}

func NewPhotoController(photo services.PhotoService) *Controller {
	return &Controller{photo: photo}
}

func NewCommentController(comment services.CommentService) *Controller {
	return &Controller{comment: comment}
}

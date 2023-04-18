package services

import "mygram/repositories"

type Service struct {
	user    repositories.UserInterface
	socmed  repositories.SocialMediaInterface
	photo   repositories.PhotoInterface
	comment repositories.CommentInterface
}

type UserInterface interface {
	UserService
}

type SocialMediaInterface interface {
	SocialMediaService
}

type PhotoInterface interface {
	PhotoService
}

type CommentInterface interface {
	CommentService
}

func NewUserService(user repositories.UserInterface) *Service {
	return &Service{user: user}
}

func NewSocialMediaService(socmed repositories.SocialMediaInterface) *Service {
	return &Service{socmed: socmed}
}

func NewPhotoService(photo repositories.PhotoInterface) *Service {
	return &Service{photo: photo}
}

func NewCommentService(comment repositories.CommentInterface) *Service {
	return &Service{comment: comment}
}

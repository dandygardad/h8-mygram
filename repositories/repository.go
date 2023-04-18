package repositories

import "gorm.io/gorm"

type Repo struct {
	gorm *gorm.DB
}

type UserInterface interface {
	UserRepository
}

type SocialMediaInterface interface {
	SocialMediaRepository
}

type PhotoInterface interface {
	PhotoRepository
}

type CommentInterface interface {
	CommentRepository
}

func NewUserRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}

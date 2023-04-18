package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"time"
)

type SocialMedia struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"type:varchar(255)"`
	SocialMediaURL string         `json:"social_media_url" gorm:"type:varchar(255)"`
	UserID         uint           `json:"user_id" gorm:"foreignKey:User"`
	User           User           `json:"user"`
	CreatedAt      time.Time      `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"type:timestamp without time zone"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}

func (s SocialMedia) Validation() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required.Error("tidak boleh kosong")),
		validation.Field(&s.SocialMediaURL, validation.Required.Error("tidak boleh kosong")),
	)
}

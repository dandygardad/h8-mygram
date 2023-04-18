package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"type:varchar(255)"`
	Caption   string         `json:"caption,omitempty" gorm:"type:varchar(255)"`
	PhotoURL  string         `json:"photo_url" gorm:"type:varchar(255)"`
	UserID    uint           `json:"user_id" gorm:"foreignKey:User"`
	User      User           `json:"user,omitempty"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp without time zone"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (p Photo) Validation() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required.Error("tidak boleh kosong")),
		validation.Field(&p.PhotoURL, validation.Required.Error("tidak boleh kosong")),
	)
}

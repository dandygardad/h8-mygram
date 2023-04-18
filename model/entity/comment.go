package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"foreignKey:User"`
	User      User           `json:"user"`
	PhotoID   uint           `json:"photo_id" gorm:"foreignKey:Photo"`
	Photo     Photo          `json:"photo"`
	Message   string         `json:"message" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp without time zone"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (c Comment) Validation() error {
	return validation.ValidateStruct(&c, validation.Field(&c.Message, validation.Required.Error("tidak boleh kosong")))
}

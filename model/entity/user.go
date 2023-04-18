package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"type:varchar(255);unique"`
	Email     string         `json:"email" gorm:"type:varchar(255);unique"`
	Password  string         `json:"-" gorm:"type:varchar(255)"`
	Age       int            `json:"age" gorm:"type:int"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp without time zone"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mygram/model/entity"
	"os"
)

type Gorm struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string

	DB *gorm.DB
}

type GormDB struct {
	*Gorm
}

var (
	NewGorm *GormDB
)

func InitGorm() error {
	NewGorm = new(GormDB)
	NewGorm.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}

	err := NewGorm.Gorm.OpenConnection()
	if err != nil {
		return err
	}
	return nil
}

func (g *Gorm) OpenConnection() error {
	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", g.Host, g.Username, g.Password, g.DBName, g.Port)
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return err
	}

	g.DB = db

	err = db.Debug().AutoMigrate(&entity.User{}, &entity.SocialMedia{}, &entity.Comment{}, &entity.Photo{})
	if err != nil {
		return err
	}
	return nil
}

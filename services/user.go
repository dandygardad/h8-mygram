package services

import (
	"errors"
	"gorm.io/gorm"
	"mygram/helper"
	"mygram/model/entity"
)

type UserService interface {
	Create(user entity.User) (entity.User, error)
	Login(user entity.User) (entity.User, error)
}

func (s *Service) Create(user entity.User) (entity.User, error) {
	// Buat validasi untuk username
	err := s.user.GetUsername(user.Username)
	if err != nil {
		if !(errors.Is(err, gorm.ErrRecordNotFound)) {
			return entity.User{}, errors.New("username_already_exist")
		}
	}

	// Buat validasi untuk email
	err = s.user.GetEmail(user.Email)
	if err != nil {
		if !(errors.Is(err, gorm.ErrRecordNotFound)) {
			return entity.User{}, errors.New("email_already_exist")
		}
	}

	hashedPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return entity.User{}, err
	}

	user.Password = hashedPass

	create, err := s.user.Create(user)
	if err != nil {
		return entity.User{}, err
	}

	create.Password = ""
	
	return create, nil
}

func (s *Service) Login(user entity.User) (entity.User, error) {
	result, err := s.user.GetOne(user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("user_not_exists")
		}
		return entity.User{}, err
	}

	// Compare password
	resultPass := helper.ComparePassword(result.Password, user.Password)
	if !resultPass {
		return entity.User{}, errors.New("wrong_password")
	}

	result.Password = ""

	return result, nil
}

package repositories

import (
	"errors"
	"mygram/model/entity"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	GetOne(user entity.User) (entity.User, error)
	GetEmail(email string) error
	GetUsername(username string) error
}

func (r *Repo) Create(user entity.User) (entity.User, error) {
	var createdUser entity.User
	err := r.gorm.Create(&user).Scan(&createdUser).Error
	if err != nil {
		return entity.User{}, err
	}
	return createdUser, nil
}

func (r *Repo) GetOne(user entity.User) (entity.User, error) {
	var resultUser entity.User
	err := r.gorm.Take(&resultUser, "username = ?", user.Username).Error
	if err != nil {
		return entity.User{}, err
	}
	return resultUser, nil
}

// GetEmail : Used for validation register
func (r *Repo) GetEmail(email string) error {
	err := r.gorm.Take(&entity.User{}, "email = ?", email).Error
	if err != nil {
		return err
	}
	return errors.New("exist")
}

// GetUsername : Used for validation register
func (r *Repo) GetUsername(username string) error {
	err := r.gorm.Take(&entity.User{}, "username = ?", username).Error
	if err != nil {
		return err
	}
	return errors.New("exist")
}

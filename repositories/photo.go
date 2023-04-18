package repositories

import (
	"gorm.io/gorm"
	"mygram/model/entity"
)

type PhotoRepository interface {
	GetAllPhoto() ([]entity.Photo, error)
	GetOnePhoto(inputPhoto entity.Photo) (entity.Photo, error)
	CreatePhoto(inputPhoto entity.Photo) (entity.Photo, error)
	UpdatePhoto(inputPhoto entity.Photo, id int) (entity.Photo, error)
	DeletePhoto(inputPhoto entity.Photo, id int) error
}

func (r *Repo) GetAllPhoto() ([]entity.Photo, error) {
	var photos []entity.Photo
	err := r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age, created_at, updated_at")
	}).Find(&photos).Error
	if err != nil {
		return []entity.Photo{}, err
	}
	return photos, nil
}

func (r *Repo) GetOnePhoto(inputPhoto entity.Photo) (entity.Photo, error) {
	var photo entity.Photo
	err := r.gorm.Preload("User").Take(&photo, "id = ?", inputPhoto.ID).Error
	if err != nil {
		return entity.Photo{}, err
	}
	photo.User.Password = ""
	return photo, nil
}

func (r *Repo) CreatePhoto(inputPhoto entity.Photo) (entity.Photo, error) {
	var result entity.Photo

	// Create
	err := r.gorm.Create(&inputPhoto).Scan(&result).Error
	if err != nil {
		return entity.Photo{}, err
	}

	// Return but preload with user
	err = r.gorm.Preload("User").Where("user_id = ?", result.UserID).First(&result).Error
	if err != nil {
		return entity.Photo{}, err
	}
	result.User.Password = "" // Remove password

	return result, nil
}

func (r *Repo) UpdatePhoto(inputPhoto entity.Photo, id int) (entity.Photo, error) {
	// Update
	var updated entity.Photo
	err := r.gorm.Preload("User").Where("user_id = ?", inputPhoto.UserID).Where("id = ?", id).Updates(&inputPhoto).Scan(&updated).Error
	if err != nil {
		return entity.Photo{}, err
	}

	// Return but preload with user
	err = r.gorm.Preload("User").Where("user_id = ?", inputPhoto.UserID).Where("id = ?", id).First(&updated).Error
	if err != nil {
		return entity.Photo{}, err
	}
	updated.User.Password = "" // Remove password

	return updated, nil
}

func (r *Repo) DeletePhoto(inputPhoto entity.Photo, id int) error {
	// Check if exist
	var checkExist []entity.Photo
	err := r.gorm.Preload("User").Where("id = ?", id).Where("user_id = ?", inputPhoto.UserID).Find(&checkExist).Error
	if err != nil {
		return err
	}
	if len(checkExist) == 0 {
		return gorm.ErrRecordNotFound
	}

	err = r.gorm.Preload("User").Where("id = ?", id).Where("user_id = ?", inputPhoto.UserID).Delete(&entity.Photo{}).Error
	if err != nil {
		return err
	}
	return nil
}

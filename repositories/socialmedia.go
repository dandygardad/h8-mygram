package repositories

import (
	"gorm.io/gorm"
	"mygram/model/entity"
)

type SocialMediaRepository interface {
	GetAllSocialMedia() ([]entity.SocialMedia, error)
	GetOneSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error)
	CreateSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error)
	UpdateSocialMedia(user entity.SocialMedia, id int) (entity.SocialMedia, error)
	DeleteSocialMedia(id int) error
}

func (r *Repo) GetAllSocialMedia() ([]entity.SocialMedia, error) {
	var socmed []entity.SocialMedia
	err := r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age, created_at, updated_at")
	}).Find(&socmed).Error
	if err != nil {
		return []entity.SocialMedia{}, err
	}
	return socmed, nil
}

func (r *Repo) GetOneSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error) {
	var socmed entity.SocialMedia
	err := r.gorm.Preload("User").Take(&socmed, "id = ?", user.ID).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	socmed.User.Password = ""
	return socmed, nil
}

func (r *Repo) CreateSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error) {
	var result entity.SocialMedia
	var checkExist []entity.SocialMedia

	// Check if exist
	errCheck := r.gorm.Preload("User").Where("user_id = ?", user.UserID).Find(&checkExist).Error
	if errCheck != nil {
		return entity.SocialMedia{}, errCheck
	}
	if len(checkExist) > 0 {
		return entity.SocialMedia{}, gorm.ErrDuplicatedKey
	}

	// Create
	err := r.gorm.Create(&user).Scan(&result).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}

	// Return but preload with user
	err = r.gorm.Preload("User").Where("user_id = ?", result.UserID).First(&result).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	result.User.Password = "" // Remove password

	return result, nil
}

func (r *Repo) UpdateSocialMedia(user entity.SocialMedia, id int) (entity.SocialMedia, error) {
	var updated entity.SocialMedia
	err := r.gorm.Preload("User").Where("user_id = ?", id).Updates(&user).Scan(&updated).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}

	// Return but preload with user
	err = r.gorm.Preload("User").Where("user_id = ?", updated.UserID).First(&updated).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	updated.User.Password = "" // Remove password

	return updated, nil
}

func (r *Repo) DeleteSocialMedia(id int) error {
	err := r.gorm.Preload("User").Where("user_id = ?", id).Delete(&entity.SocialMedia{}).Error
	if err != nil {
		return err
	}
	return nil
}

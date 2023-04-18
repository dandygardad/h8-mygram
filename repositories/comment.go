package repositories

import (
	"errors"
	"gorm.io/gorm"
	"mygram/model/entity"
)

type CommentRepository interface {
	CheckJWTComment(id int) error
	GetAllComment(id int) ([]entity.Comment, error)
	GetOneComment(inputComment entity.Comment) (entity.Comment, error)
	CreateComment(inputComment entity.Comment) (entity.Comment, error)
	UpdateComment(inputComment entity.Comment, id int) (entity.Comment, error)
	DeleteComment(inputComment entity.Comment, id int) error
}

func (r *Repo) CheckJWTComment(id int) error {
	var exist []entity.User
	err := r.gorm.Where("id = ?", id).Find(&exist).Error
	if err != nil {
		return err
	}
	if len(exist) == 0 {
		return errors.New("not_founs")
	}
	return nil
}

func (r *Repo) GetAllComment(id int) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.gorm.Preload("User").Preload("Photo").Preload("Photo.User").Where("photo_id = ?", id).Find(&comments).Error
	if err != nil {
		return []entity.Comment{}, err
	}
	return comments, nil
}

func (r *Repo) GetOneComment(inputComment entity.Comment) (entity.Comment, error) {
	var comment entity.Comment
	err := r.gorm.Preload("User").Preload("Photo").Preload("Photo.User").Take(&comment, "id = ?", inputComment.Id).Error
	if err != nil {
		return entity.Comment{}, err
	}
	comment.User.Password = ""
	return comment, nil
}

func (r *Repo) CreateComment(inputComment entity.Comment) (entity.Comment, error) {
	// Check if photo deleted
	var checkExist []entity.Photo
	_ = r.gorm.Preload("User").Where("id", inputComment.PhotoID).Find(&checkExist).Error
	if len(checkExist) == 0 {
		return entity.Comment{}, errors.New("photo_not_found")
	}

	// Create
	var result entity.Comment
	err := r.gorm.Create(&inputComment).Scan(&result).Error
	if err != nil {
		return entity.Comment{}, err
	}

	// Return but preload with user
	err = r.gorm.Preload("User").Preload("Photo").Preload("Photo.User").Where("photo_id = ?", result.PhotoID).First(&result).Error
	if err != nil {
		return entity.Comment{}, err
	}

	return result, nil
}

func (r *Repo) UpdateComment(inputComment entity.Comment, id int) (entity.Comment, error) {
	// Update
	var updated entity.Comment
	err := r.gorm.Preload("User").Preload("Photo").Preload("Photo.User").Where("user_id = ?", inputComment.UserID).Where("id = ?", id).Updates(&inputComment).Scan(&updated).Error
	if err != nil {
		return entity.Comment{}, err
	}

	// Return but preload with user
	err = r.gorm.Preload("User").Preload("Photo").Preload("Photo.User").Where("user_id = ?", inputComment.UserID).Where("id = ?", id).First(&updated).Error
	if err != nil {
		return entity.Comment{}, err
	}
	updated.User.Password = "" // Remove password

	return updated, nil
}

func (r *Repo) DeleteComment(inputComment entity.Comment, id int) error {
	// Check if exist
	var checkExist []entity.Comment
	err := r.gorm.Preload("User").Where("id = ?", id).Where("user_id = ?", inputComment.UserID).Find(&checkExist).Error
	if err != nil {
		return err
	}
	if len(checkExist) == 0 {
		return gorm.ErrRecordNotFound
	}

	err = r.gorm.Preload("User").Where("id = ?", id).Where("user_id = ?", inputComment.UserID).Delete(&entity.Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

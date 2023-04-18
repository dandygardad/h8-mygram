package services

import (
	"errors"
	"gorm.io/gorm"
	"mygram/model/entity"
)

type CommentService interface {
	GetAllComment(id int) ([]entity.Comment, error)
	GetOneComment(inputComment entity.Comment) (entity.Comment, error)
	CreateComment(inputComment entity.Comment) (entity.Comment, error)
	UpdateComment(inputComment entity.Comment, id int) (entity.Comment, error)
	DeleteComment(inputComment entity.Comment, id int) error
}

func (s *Service) GetAllComment(id int) ([]entity.Comment, error) {
	results, err := s.comment.GetAllComment(id)
	if err != nil {
		return []entity.Comment{}, err
	}

	if len(results) == 0 {
		return []entity.Comment{}, errors.New("photo_not_found")
	}

	return results, nil
}

func (s *Service) GetOneComment(inputComment entity.Comment) (entity.Comment, error) {
	result, err := s.comment.GetOneComment(inputComment)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Comment{}, errors.New("not_found")
		}
		return entity.Comment{}, err
	}
	return result, nil
}

func (s *Service) CreateComment(inputComment entity.Comment) (entity.Comment, error) {
	err := s.comment.CheckJWTComment(int(inputComment.UserID))
	if err != nil {
		return entity.Comment{}, errors.New("not_found")
	}

	created, err := s.comment.CreateComment(inputComment)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Comment{}, errors.New("already_exist")
		}
		return entity.Comment{}, err
	}
	return created, nil
}

func (s *Service) UpdateComment(inputComment entity.Comment, id int) (entity.Comment, error) {
	result, err := s.comment.UpdateComment(inputComment, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Comment{}, errors.New("not_found")
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Comment{}, errors.New("already_exist")
		}
		return entity.Comment{}, err
	}
	return result, nil
}

func (s *Service) DeleteComment(inputComment entity.Comment, id int) error {
	err := s.comment.DeleteComment(inputComment, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not_found")
		}
		return err
	}
	return nil
}

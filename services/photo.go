package services

import (
	"errors"
	"gorm.io/gorm"
	"mygram/model/entity"
)

type PhotoService interface {
	GetAllPhoto() ([]entity.Photo, error)
	GetOnePhoto(inputUser entity.Photo) (entity.Photo, error)
	CreatePhoto(inputUser entity.Photo) (entity.Photo, error)
	UpdatePhoto(inputUser entity.Photo, id int) (entity.Photo, error)
	DeletePhoto(inputUser entity.Photo, id int) error
}

func (s *Service) GetAllPhoto() ([]entity.Photo, error) {
	results, err := s.photo.GetAllPhoto()
	if err != nil {
		return []entity.Photo{}, err
	}

	if len(results) == 0 {
		return []entity.Photo{}, errors.New("no_data")
	}

	return results, nil
}

func (s *Service) GetOnePhoto(inputUser entity.Photo) (entity.Photo, error) {
	result, err := s.photo.GetOnePhoto(inputUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Photo{}, errors.New("not_found")
		}
		return entity.Photo{}, err
	}
	return result, nil
}

func (s *Service) CreatePhoto(inputUser entity.Photo) (entity.Photo, error) {
	created, err := s.photo.CreatePhoto(inputUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Photo{}, errors.New("already_exist")
		}
		return entity.Photo{}, err
	}
	return created, nil
}

func (s *Service) UpdatePhoto(inputUser entity.Photo, id int) (entity.Photo, error) {
	result, err := s.photo.UpdatePhoto(inputUser, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Photo{}, errors.New("not_found")
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Photo{}, errors.New("already_exist")
		}
		return entity.Photo{}, err
	}
	return result, nil
}

func (s *Service) DeletePhoto(inputUser entity.Photo, id int) error {
	err := s.photo.DeletePhoto(inputUser, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not_found")
		}
		return err
	}
	return nil
}

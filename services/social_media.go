package services

import (
	"errors"
	"gorm.io/gorm"
	"mygram/model/entity"
)

type SocialMediaService interface {
	GetAllSocialMedia() ([]entity.SocialMedia, error)
	GetOneSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error)
	CreateSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error)
	UpdateSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error)
	DeleteSocialMedia(id int) error
}

func (s *Service) GetAllSocialMedia() ([]entity.SocialMedia, error) {
	results, err := s.socmed.GetAllSocialMedia()
	if err != nil {
		return []entity.SocialMedia{}, err
	}

	if len(results) == 0 {
		return []entity.SocialMedia{}, errors.New("no_data")
	}

	return results, nil
}

func (s *Service) GetOneSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error) {
	result, err := s.socmed.GetOneSocialMedia(user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SocialMedia{}, errors.New("not_found")
		}
		return entity.SocialMedia{}, err
	}
	return result, nil
}

func (s *Service) CreateSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error) {
	created, err := s.socmed.CreateSocialMedia(user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.SocialMedia{}, errors.New("already_exist")
		}
		return entity.SocialMedia{}, err
	}
	return created, nil
}

func (s *Service) UpdateSocialMedia(user entity.SocialMedia) (entity.SocialMedia, error) {
	result, err := s.socmed.UpdateSocialMedia(user, int(user.UserID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SocialMedia{}, errors.New("not_found")
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.SocialMedia{}, errors.New("already_exist")
		}
		return entity.SocialMedia{}, err
	}
	return result, nil
}

func (s *Service) DeleteSocialMedia(id int) error {
	err := s.socmed.DeleteSocialMedia(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not_found")
		}
		return err
	}
	return nil
}

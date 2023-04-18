package web

import validation "github.com/go-ozzo/ozzo-validation"

type SocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (c SocialMediaRequest) Validation() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error("tidak boleh kosong")),
		validation.Field(&c.SocialMediaURL, validation.Required.Error("tidak boleh kosong")),
	)
}

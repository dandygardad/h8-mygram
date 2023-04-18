package web

import validation "github.com/go-ozzo/ozzo-validation"

type PhotoRequest struct {
	Title    string `json:"title"`
	PhotoURL string `json:"photo_url"`
	Caption  string `json:"caption"`
}

func (c PhotoRequest) Validation() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required.Error("tidak boleh kosong")),
		validation.Field(&c.PhotoURL, validation.Required.Error("tidak boleh kosong")),
	)
}

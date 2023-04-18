package web

import validation "github.com/go-ozzo/ozzo-validation"

type CommentRequest struct {
	Message string `json:"message" gorm:"type:text"`
}

func (c CommentRequest) Validation() error {
	return validation.ValidateStruct(&c, validation.Field(&c.Message, validation.Required.Error("tidak boleh kosong")))
}

package web

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u RegisterRequest) Validation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required.Error("tidak boleh kosong"), is.Email.Error("email tidak valid")),
		validation.Field(&u.Username, validation.Required.Error("tidak boleh kosong")),
		validation.Field(&u.Password, validation.Required.Error("tidak boleh kosong"), validation.Length(6, 255).Error("minimal harus 6 karakter")),
		validation.Field(&u.Age, validation.Required.Error("tidak boleh kosong"), validation.Min(8).Error("minimal umur harus 8 tahun")),
	)
}

func (u LoginRequest) Validation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required.Error("tidak boleh kosong")),
		validation.Field(&u.Password, validation.Required.Error("tidak boleh kosong"), validation.Length(6, 255).Error("minimal harus 6 karakter")),
	)
}

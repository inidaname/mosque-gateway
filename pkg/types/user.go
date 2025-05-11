package types

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

func (p CreateUserPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.FullName, validation.Required),
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.Password, validation.Required),
		validation.Field(&p.Phone, validation.Required),
	)
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (p LoginUserPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.Password, validation.Required),
	)
}

type ForgotPasswordPayload struct {
	Email string `json:"email" validate:"required,email"`
}

func (p ForgotPasswordPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required, is.Email),
	)
}

type ValidatePasswordToken struct {
	Token string `json:"token" validate:"required"`
}

func (p ValidatePasswordToken) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Token, validation.Required),
	)
}

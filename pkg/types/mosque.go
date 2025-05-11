package types

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateMosquePayload struct {
	Name       string    `json:"name" validate:"required"`
	Address    string    `json:"address" validate:"required"`
	EidTime    time.Time `json:"eid_time" validate:"required"`
	JummahTime time.Time `json:"jummah_time" validate:"required"`
	Lat        float64   `json:"lat" validate:"required"`
	Lng        float64   `json:"lng" validate:"required"`
}

func (p CreateMosquePayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Address, validation.Required),
		validation.Field(&p.EidTime, validation.Required),
		validation.Field(&p.JummahTime, validation.Required),
		validation.Field(&p.Lat, validation.Required),
		validation.Field(&p.Lng, validation.Required),
	)
}

type UpdateMosqueMosquePayload struct{}

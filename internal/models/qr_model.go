package model

import (
	"github.com/go-playground/validator/v10"
)

// QRCodeInput represents the input data for generating a QR code
type QRCodeInput struct {
	Content   string `json:"content" validate:"required"`            // Content to encode in the QR code
	Scheduler bool   `json:"scheduler" validate:"omitempty,boolean"` // New field to enable/disable scheduler
}

// Validate function to validate QRCodeInput
func (input *QRCodeInput) Validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

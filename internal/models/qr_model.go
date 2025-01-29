package model

import (
	"github.com/go-playground/validator/v10"
)

// QRCodeInput represents the input data for generating a QR code
type QRCodeInput struct {
	Data string `json:"data" validate:"required,url"` // Validate that Data is a required URL
}

// Validate function to validate QRCodeInput
func (input *QRCodeInput) Validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

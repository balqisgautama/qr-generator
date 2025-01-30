package model

import (
	"github.com/go-playground/validator/v10"
)

// QRCodeInput represents the input data for generating a QR code
type QRCodeInput struct {
	Content             string `json:"content" validate:"required"`                             // Content to encode in the QR code
	IsSchedulerDeleteOn bool   `json:"is_scheduler_delete_on" validate:"omitempty,boolean"`     // New field to enable/disable scheduler
	IsUsingCustomLogo   bool   `json:"is_using_custom_logo" validate:"omitempty,boolean"`       // New field to enable/disable custom logo
	FileName            string `json:"file_name" validate:"required_if=IsUsingCustomLogo true"` // New field to specify the file name
}

// Validate function to validate QRCodeInput
func (input *QRCodeInput) Validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

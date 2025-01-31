package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadLogoService handles logo uploads
type UploadLogoService struct {
	logoDir string
}

// NewUploadLogoService initializes a new UploadLogoService
func NewUploadLogoService(logoDir string) (*UploadLogoService, error) {
	if err := os.MkdirAll(logoDir, os.ModePerm); err != nil {
		return nil, err
	}
	return &UploadLogoService{
		logoDir: logoDir,
	}, nil
}

// UploadLogo handles the logo upload
func (s *UploadLogoService) UploadLogo(c *gin.Context) (string, string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", "", err
	}

	// Validate the uploaded file type
	if err := s.ValidateImage(*file); err != nil {
		return "", "", err
	}

	// Save the uploaded logo to a temporary location
	filename := fmt.Sprintf("logo_%s.png", time.Now().Format("20060102150405"))
	tempLogoPath := filepath.Join(s.logoDir, filename)
	if err := c.SaveUploadedFile(file, tempLogoPath); err != nil {
		return "", "", err
	}

	return filename, tempLogoPath, nil // Return the path of the saved logo
}

// ValidateImage checks if the uploaded file is an image
func (s *UploadLogoService) ValidateImage(file multipart.FileHeader) error {
	// Get the file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	// Allowed image extensions
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}

	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return nil
		}
	}
	return errors.New("only image files (png, jpg, jpeg) are allowed")
}

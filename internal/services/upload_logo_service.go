package service

import (
	"fmt"
	"os"
	"path/filepath"
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

	// Save the uploaded logo to a temporary location
	filename := fmt.Sprintf("logo_%s.png", time.Now().Format("20060102150405"))
	tempLogoPath := filepath.Join(s.logoDir, filename)
	if err := c.SaveUploadedFile(file, tempLogoPath); err != nil {
		return "", "", err
	}

	return filename, tempLogoPath, nil // Return the path of the saved logo
}

package service

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
)

// QRCodeService handles QR code generation and cleanup
type QRCodeService struct {
	qrCodeDir           string
	logoDir             string
	expirationTime      time.Duration
	IsSchedulerDeleteOn bool
	mu                  sync.Mutex
}

// NewQRCodeService initializes a new QRCodeService
func NewQRCodeService(qrCodeDir string, logoDir string) (*QRCodeService, error) {
	if err := os.MkdirAll(qrCodeDir, os.ModePerm); err != nil {
		return nil, err
	}
	return &QRCodeService{
		qrCodeDir:           qrCodeDir,
		logoDir:             logoDir,
		expirationTime:      60 * time.Second, // Default expiration time
		IsSchedulerDeleteOn: false,            // Default scheduler on
	}, nil
}

// SetExpiration sets the expiration time for QR codes
func (s *QRCodeService) SetExpiration(duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.expirationTime = duration
}

// GenerateQRCode generates a QR code and saves it to a file
func (s *QRCodeService) GenerateQRCode(content string) (string, string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filename := fmt.Sprintf("qr_%s.png", time.Now().Format("20060102150405"))
	filePath := filepath.Join(fmt.Sprintf(`./%s`, s.qrCodeDir), filename)

	err := qrcode.WriteFile(content, qrcode.Medium, 256, filePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Schedule the file to be deleted if the scheduler is enabled
	if s.IsSchedulerDeleteOn {
		go func() {
			time.Sleep(s.expirationTime)
			if err := os.Remove(filePath); err != nil {
				log.Printf("Failed to delete QR code file: %v\n", err)
			}
		}()
	}

	return filename, filePath, nil
}

// GetQRCodeDir returns the directory where QR codes are stored
func (s *QRCodeService) GetQRCodeDir() string {
	return s.qrCodeDir
}

func (s *QRCodeService) GetLogoDir() string {
	return s.logoDir
}

// AddLogoToQRCode overlays a logo on the generated QR code with a border
func (s *QRCodeService) AddLogoToQRCode(qrCodePath string, logoPath string) (string, string, error) {
	// Load the QR code
	qrCodeFile, err := os.Open(qrCodePath)
	if err != nil {
		return "", "", err
	}
	defer qrCodeFile.Close()

	qrCodeImg, err := png.Decode(qrCodeFile)
	if err != nil {
		return "", "", err
	}

	// Load the logo
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return "", "", err
	}
	defer logoFile.Close()

	logoImg, err := png.Decode(logoFile)
	if err != nil {
		return "", "", err
	}

	// Resize the logo to fit in the center of the QR code
	logoSize := 50 // Size of the logo
	logoImg = imaging.Resize(logoImg, logoSize, logoSize, imaging.Lanczos)

	// Create a new image to hold the QR code with the logo
	qrWithLogo := image.NewRGBA(qrCodeImg.Bounds())
	draw.Draw(qrWithLogo, qrCodeImg.Bounds(), qrCodeImg, image.Point{0, 0}, draw.Over)

	// Calculate the offset to center the logo
	offset := (qrWithLogo.Bounds().Size().X - logoImg.Bounds().Size().X) / 2

	// Draw a border around the logo
	borderSize := 5 // Size of the border
	borderRect := image.Rect(offset-borderSize, offset-borderSize, offset+logoImg.Bounds().Size().X+borderSize, offset+logoImg.Bounds().Size().Y+borderSize)

	// Create a white border
	borderColor := color.White
	draw.Draw(qrWithLogo, borderRect, &image.Uniform{borderColor}, image.Point{}, draw.Src)

	// Draw the logo on top of the border
	draw.Draw(qrWithLogo, logoImg.Bounds().Add(image.Point{offset, offset}), logoImg, image.Point{0, 0}, draw.Over)

	// Save the QR code with logo to a file
	finalFilename := fmt.Sprintf("qr_with_logo_%s.png", time.Now().Format("20060102150405"))
	finalFilePath := filepath.Join(s.qrCodeDir, finalFilename)
	finalFile, err := os.Create(finalFilePath)
	if err != nil {
		return "", "", err
	}
	defer finalFile.Close()

	if err := png.Encode(finalFile, qrWithLogo); err != nil {
		return "", "", err
	}

	return finalFilename, finalFilePath, nil // Return just the filename for download
}

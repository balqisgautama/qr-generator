package service

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/skip2/go-qrcode"
)

// QRCodeService handles QR code generation and cleanup
type QRCodeService struct {
	qrCodeDir      string
	expirationTime time.Duration
	schedulerOn    bool
	mu             sync.Mutex
}

// NewQRCodeService initializes a new QRCodeService
func NewQRCodeService(qrCodeDir string) (*QRCodeService, error) {
	if err := os.MkdirAll(qrCodeDir, os.ModePerm); err != nil {
		return nil, err
	}
	return &QRCodeService{
		qrCodeDir:      qrCodeDir,
		expirationTime: 5 * time.Minute, // Default expiration time
		schedulerOn:    true,            // Default scheduler on
	}, nil
}

// SetExpiration sets the expiration time for QR codes
func (s *QRCodeService) SetExpiration(duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.expirationTime = duration
}

// GenerateQRCode generates a QR code and saves it to a file
func (s *QRCodeService) GenerateQRCode(content string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filename := fmt.Sprintf("%s.png", time.Now().Format("20060102150405"))
	filePath := filepath.Join(s.qrCodeDir, filename)

	err := qrcode.WriteFile(content, qrcode.Medium, 256, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Schedule the file to be deleted if the scheduler is enabled
	if s.schedulerOn {
		go func() {
			time.Sleep(s.expirationTime)
			if err := os.Remove(filePath); err != nil {
				fmt.Printf("Failed to delete QR code file: %v\n", err)
			}
		}()
	}

	return filePath, nil
}

// GetQRCodeDir returns the directory where QR codes are stored
func (s *QRCodeService) GetQRCodeDir() string {
	return s.qrCodeDir
}

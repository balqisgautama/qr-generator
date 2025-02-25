package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	model "qr-generator/internal/models"
	service "qr-generator/internal/services"

	"github.com/gin-gonic/gin"
)

// QRCodeHandler handles HTTP requests for QR code generation
type QRCodeHandler struct {
	serviceQR *service.QRCodeService
}

// NewQRCodeHandler initializes a new QRCodeHandler
func NewQRCodeHandler(service *service.QRCodeService) *QRCodeHandler {
	return &QRCodeHandler{
		serviceQR: service,
	}
}

// GenerateQRCodeHandler handles the POST request to generate a QR code
func (h *QRCodeHandler) GenerateQRCodeHandler(c *gin.Context) {
	var input model.QRCodeInput

	// Bind JSON and validate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.serviceQR.IsSchedulerDeleteOn = input.IsSchedulerDeleteOn
	fileName, filePath, err := h.serviceQR.GenerateQRCode(input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if condition := input.IsUsingCustomLogo; condition {
		// Add logo to the QR code
		logoPath := fmt.Sprintf(`./%s/%s`, h.serviceQR.GetLogoDir(), input.FileName)
		finalFileName, finalFilePath, err := h.serviceQR.AddLogoToQRCode(filePath, logoPath)
		if err != nil {
			if condition := os.IsNotExist(err); condition {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("logo file not found: %w", err).Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fileName = finalFileName
		filePath = finalFilePath
	}

	c.JSON(http.StatusOK, gin.H{"message": "QR code generated successfully", "file_name": fileName, "file_path": filePath})
}

// DownloadQRCodeHandler handles the GET request to download the QR code image
func (h *QRCodeHandler) DownloadQRCodeHandler(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(h.serviceQR.GetQRCodeDir(), filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Set the headers to trigger a download
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "image/png")
	c.File(filePath)
}

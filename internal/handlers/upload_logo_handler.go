package handler

import (
	"net/http"
	service "qr-generator/internal/services"

	"github.com/gin-gonic/gin"
)

// LogoHandler handles logo uploads
type UploadLogoHandler struct {
	serviceUploadLogo *service.UploadLogoService
}

// NewUploadLogoHandler initializes a new UploadLogoHandler
func NewUploadLogoHandler(service *service.UploadLogoService) *UploadLogoHandler {
	return &UploadLogoHandler{
		serviceUploadLogo: service,
	}
}

// UploadImageHandler handles the POST request to upload a logo
func (h *UploadLogoHandler) UploadLogoHandler(c *gin.Context) {
	filename, path, err := h.serviceUploadLogo.UploadLogo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logo uploaded successfully", "filename": filename, "path": path})
}

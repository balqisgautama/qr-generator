package main

import (
	"log"
	handler "qr-generator/internal/handlers"
	service "qr-generator/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize QRCodeService
	qrCodeDir := "./assets/qr_codes"
	logoDir := "./assets/logos"

	qrService, err := service.NewQRCodeService(qrCodeDir, logoDir)
	if err != nil {
		log.Fatalf("Failed to initialize QRCodeService: %v\n", err)
	}

	// Initialize UploadLogoService
	uploadLogoService, err := service.NewUploadLogoService(logoDir)
	if err != nil {
		log.Fatalf("Failed to initialize UploadService: %v\n", err)
	}

	// Initialize QRCodeHandler
	qrHandler := handler.NewQRCodeHandler(qrService)

	// Initialize UploadLogoHandler
	uploadHandler := handler.NewUploadLogoHandler(uploadLogoService)

	// Initialize Gin router
	r := gin.Default()

	// Define the POST endpoint for generating QR codes
	r.POST("/generate-qr", qrHandler.GenerateQRCodeHandler)

	// Define the GET endpoint for downloading QR codes
	r.GET("/download-qr/:filename", qrHandler.DownloadQRCodeHandler)

	// Define the POST endpoint for uploading logos
	r.POST("/upload-logo", uploadHandler.UploadLogoHandler)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}

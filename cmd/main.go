package main

import (
	"log"
	handler "qr-generator/internal/handlers"
	service "qr-generator/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize QRCodeService
	qrCodeDir := "./assets"
	qrService, err := service.NewQRCodeService(qrCodeDir)
	if err != nil {
		log.Fatalf("Failed to initialize QRCodeService: %v\n", err)
	}

	// Initialize QRCodeHandler
	qrHandler := handler.NewQRCodeHandler(qrService)

	// Initialize Gin router
	r := gin.Default()

	// Define the POST endpoint for generating QR codes
	r.POST("/generate-qr", qrHandler.GenerateQRCodeHandler)

	// Define the GET endpoint for downloading QR codes
	r.GET("/download-qr/:filename", qrHandler.DownloadQRCodeHandler)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}

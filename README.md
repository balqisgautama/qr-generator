# QR Code Generator API

## Overview
The QR Code Generator API is a simple RESTful API built with Go that allows users to generate QR codes based on input data (URLs) and download the generated QR code images. The API supports automatic deletion of QR codes after a specified expiration time.

## Features
- Generate QR codes from user-provided URLs.
- Download generated QR code images.
- Input validation to ensure URLs are valid.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Clone the Repository](#clone-the-repository)
  - [Install Dependencies](#install-dependencies)
- [Usage](#usage)
  - [Run the API Server](#run-the-api-server)
  - [Generate a QR Code](#generate-a-qr-code)
  - [Download a QR Code](#download-a-qr-code)
- [API Endpoints](#api-endpoints)
  - [Generate QR Code](#generate-qr-code)
  - [Download QR Code](#download-qr-code)
- [Docker Setup](#docker-setup)
  - [Using Docker Compose](#using-docker-compose)
- [OpenAPI Specification](#openapi-specification)
- [Contributing](#contributing)

## Installation

### Prerequisites
- Go (version 1.16 or later)
- Git
- Docker (for Docker setup)

### Clone the Repository
```bash
git clone https://github.com/balqisgautama/qr-generator.git
cd qr-generator
```

### Install Dependencies
Run the following command to install the required Go packages:
```bash
go mod tidy
```

## Usage

### Run the API Server
To start the API server, run:
```bash
go run ./cmd/main.go
```
The server will start on `http://localhost:8080`.

### Generate a QR Code
To generate a QR code, send a POST request to the `/generate-qr` endpoint with a JSON body containing the URL to encode.

**Example Request:**
```bash
curl -X POST http://localhost:8080/generate-qr -H "Content-Type: application/json" -d '{"data": "https://example.com"}'
```

**Example Response:**
```json
{
    "message": "QR code generated successfully",
    "file_path": "assets/20250129043049.png",
}
```

### Download a QR Code
To download a generated QR code image, send a GET request to the `/download-qr/{filename}` endpoint, replacing `{filename}` with the name of the QR code image file.

**Example Request:**
```bash
curl -O http://localhost:8080/download-qr/20250129043049.png
```

## API Endpoints

### Generate QR Code
- **Endpoint:** `/generate-qr`
- **Method:** `POST`
- **Request Body:**
  - `data` (string, required): The URL to encode in the QR code (must be a valid URL).
- **Responses:**
  - `200 OK`: QR code generated successfully.
  - `400 Bad Request`: Invalid input data.
  - `500 Internal Server Error`: Error generating QR code.

### Download QR Code
- **Endpoint:** `/download-qr/{filename}`
- **Method:** `GET`
- **Path Parameter:**
  - `filename` (string, required): The name of the QR code image file to download.
- **Responses:**
  - `200 OK`: Returns the QR code image file.
  - `404 Not Found`: The specified file does not exist.

## Docker Setup

### Using Docker Compose
To deploy the QR Code Generator API using Docker, you can use the provided `Dockerfile` and `docker-compose.yml`.

1. **Build and Run the Application**:
   In the root directory of your project, run:
   ```bash
   make build
   ```

2. **Access the API**:
   The API will be accessible at `http://localhost:8080`.

3. **Stop the Application**:
   To stop the application, run:
   ```bash
   make stop clean
   ```

## OpenAPI Specification
The API is documented using OpenAPI Specification (OAS). You can find the OAS file in the `oas.yml` file in the root of the repository. This file can be used with tools like Swagger UI or Postman for interactive API documentation.

## Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.
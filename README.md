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
- Go (version 1.23 or later)
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
curl -X POST http://localhost:8080/generate-qr -H "Content-Type: application/json" -d '{"content": "https://github.com/balqisgautama/generate-qr", "is_scheduler_delete_on": true, "is_using_custom_logo": true, "file_name": "qr_20250130135140.png"}'
```

**Example Response:**
```json
{
    "file_name": "20250129142508.png",
    "file_path": "assets/qr_codes/20250129142508.png",
    "message": "QR code generated successfully"
}
```

### Download a QR Code
To download a generated QR code image, send a GET request to the `/download-qr/{filename}` endpoint, replacing `{filename}` with the name of the QR code image file.

**Example Request:**
```bash
curl -O http://localhost:8080/download-qr/20250129043049.png
```

### Upload Logo
To upload a logo, send a POST request to the /upload-logo endpoint with the logo file.

**Example Request:**
```bash
curl -X POST http://localhost:8080/upload-logo -H "Content-Type: multipart/form-data" -F "image=@/path/to/logo.png"
```

**Example Response:**
```json
{
    "file_name": "logo_20250129142541.png",
    "file_path": "assets/logos/logo_20250129142541.png",
    "message": "Logo uploaded successfully"
}
```

## API Endpoints

### Generate QR Code
- **Endpoint:** `/generate-qr`
- **Method:** `POST`
- **Request Body:**
  - `content` (string, required): The data to encode in the QR code.
  - `scheduler` (boolean, omitempty): Enable/disable scheduler. Default expiration time is a minute.
  - `is_using_custom_logo` (boolean, omitempty): Indicates if a custom logo should be used.
  - `file_name` (string, omitempty): The file name to the uploaded logo to use if `is_using_custom_logo` is true.
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

### Upload Logo
- **Endpoint:** `/upload-logo`
- **Method:** `POST`
- **Request Body:** `application/json`
  - `logo` (file, required): The logo image file to upload.
- **Responses:**
  - `200 OK`: Logo uploaded successfully.
  - `400 Bad Reuqest`: Invalid logo file.

## Docker Setup

### Using Docker Compose
To deploy the QR Code Generator API using Docker, you can use the provided `Dockerfile` and `docker-compose.yml`.

1. **Build and Run the Application**:
   In the root directory of your project, run:
   ```bash
   make run
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
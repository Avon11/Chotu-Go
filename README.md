# Chotu-Go: URL Shortener Service

Chotu-Go is a simple and efficient URL shortener service built with Go. It allows users to create short, easy-to-share URLs that redirect to longer, original URLs.

## Features

- Create short URLs from long ones
- Redirect short URLs to their original destinations
- Fast and efficient, using Redis for caching
- MongoDB for persistent storage

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.16 or higher
- Redis server
- MongoDB server
- Git (for cloning the repository)

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/Avon11/Chotu-Go.git
   ```

2. Navigate to the project directory:

   ```
   cd Chotu-Go
   ```

3. Install dependencies:

   ```
   go mod tidy
   ```

4. Set up your environment variables (see Configuration section below)

5. Build the project:
   ```
   go build
   ```

## Configuration

Create a `.env` file in the root of the project with the following variables:

```
MONGO_URI=mongodb:
```

Adjust these values according to your Redis and MongoDB configurations.

## Usage

To start the server, run:

```
go run cmd/main.go
```

The server will start on `http://localhost:8080` by default.

### API Endpoints

1. Create a short URL

   - Endpoint: `POST /post-url`
   - Request body:
     ```json
     {
       "url": "https://example.com/very/long/url/that/needs/shortening"
     }
     ```
   - Response:
     ```json
     {
       "code": 200,
       "msg": "success",
       "model": {
         "url": "https://chotu.com/06SWjw"
       }
     }
     ```

2. Redirect to original URL

   - Endpoint: `GET /get-url?code=abc123`
   - Response: Get redirect url with code `200`
     ```json
     {
       "code": 200,
       "msg": "success",
       "model": {
         "shortCode": "06SWjw",
         "url": "https://github.com/Avon11/User-management/blob/main/main.go"
       }
     }
     ```

3. Check original URL before redirecting - Add `_` at end of shortcode
   - Endpoint: `GET /get-url?code=abc123_`
   - Response: Get redirect url with code `201`
     ```json
     {
       "code": 201,
       "msg": "success",
       "model": {
         "shortCode": "06SWjw",
         "url": "https://github.com/Avon11/User-management/blob/main/main.go"
       }
     }
     ```

## Future Improvements

1. FE interface
2. QR code generation
3. Link expiration and Self destruct

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Go-Redis](https://github.com/go-redis/redis)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

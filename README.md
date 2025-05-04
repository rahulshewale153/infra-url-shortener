# infracloud-url-shortener

A simple, in-memory URL shortener service built with Go. It provides REST APIs to shorten URLs, redirect shortened links to original URLs, and view metrics of top three shortened domains.


## Features

- Shorten long URLs and get a unique short URL
- Idempotent shortening (same long URL returns same short URL)
- Redirect short URL to the original URL
- Metrics API to return top 3 most-shortened domains
- In-memory storage (no database required)
- Graceful shutdown support
- Docker support

## Prerequisites
- Docker

## Run with Docker
1. Build the Docker image:
   ```bash
   docker build -t infracloud-url-shortener .
   ```
2. Run the Docker container:
   ```bash
    docker run -p 8080:8080 infracloud-url-shortener
    ```
3. Access the service at `http://localhost:8080`


## API Endpoints
### Shorten URL
- **Endpoint**: `POST /url/shortener`
- **Request Body**:
  ```json
  {
    "url": "https://www.infracloud.io/careers"
  }
  ```

- **Response**:
  ```json
  {
    "short_url": "http://localhost:8080/abc123"
  }
  ```

### Redirect Short URL
- **Endpoint**: `GET /{short_url_id}`
- **Response**: Redirects to the original URL.


### For Top 3 Domains
- **Endpoint**: `GET /url/top-domains`
- **Response**:
  ```
  ["https://www.infracloud.io", "https://www.example.com", "https://www.google.com"]
  
  ```

## Image access link
- [Docker Hub](https://hub.docker.com/r/rahulshewale153/urlshortener)
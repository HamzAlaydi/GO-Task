# Simple HTTP API - Hello World Service

A lightweight Go web service that validates names based on their first letter and returns appropriate responses.

## Overview

This service provides a single endpoint `/hello-world` that accepts a `name` query parameter and responds based on whether the first letter falls in the first half of the English alphabet (A–M).

## Requirements

- Go 1.18 or later
- No external dependencies (uses standard library only)

## Installation

1. Clone the repository:

```bash
git clone <your-repo-url>
cd <repo-directory>
```

2. Initialize the Go module (if not already done):

```bash
go mod init hello-world-api
```

## Running the Application

Start the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

Expected log output:

```
2025/11/09 12:00:00 Server starting on port :8080
```

## Running the Tests

Run all tests:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -v -cover
```

Generate coverage report:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## API Documentation

### Endpoint: `GET /hello-world`

#### Parameters

- `name` (query parameter, required): The name to validate.

#### Response Scenarios

**Valid Name (First Letter A–M)**

```bash
curl "http://localhost:8080/hello-world?name=Alice"
```

```json
{
  "message": "Hello Alice"
}
```

Status: `200 OK`

**Invalid Name (First Letter N–Z)**

```bash
curl "http://localhost:8080/hello-world?name=Zane"
```

```json
{
  "error": "Invalid Input"
}
```

Status: `400 Bad Request`

**Missing/Empty Name**

```bash
curl "http://localhost:8080/hello-world?name="
```

```json
{
  "error": "Invalid Input"
}
```

Status: `400 Bad Request`

## Design Decisions & Assumptions

- Case insensitive validation using the first character after trimming whitespace.
- Inputs whose first character is not a letter are considered invalid.
- Only `GET` requests are supported; other methods return `405 Method Not Allowed`.
- The service listens on port `8080` by default.

## Project Structure

```
.
├── main.go      # Application entry point and HTTP handler
├── main_test.go # Comprehensive unit tests
├── go.mod       # Go module definition
└── README.md    # Project documentation
```

## Example Usage

```bash
# Valid requests
curl "http://localhost:8080/hello-world?name=Alice"
curl "http://localhost:8080/hello-world?name=bob"

# Invalid requests
curl "http://localhost:8080/hello-world?name=Nancy"
curl "http://localhost:8080/hello-world?name="
```

## Future Enhancements

- Configuration management (e.g. environment-based port).
- Structured logging with log levels.
- Health check endpoint and metrics.
- Graceful shutdown.
- Containerization with Docker and CI/CD automation.

## License

This is a coding assessment project.

# GO-Task

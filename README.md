# Simple HTTP API - Hello World Service

A lightweight Go web service that validates names based on their first letter and returns appropriate responses.

## Overview

This service provides a single endpoint `/hello-world` that accepts a name query parameter and responds based on whether the first letter falls in the first half of the English alphabet (A-M).

## Requirements

- Go 1.18 or later
- No external dependencies (uses standard library only)

## Installation

1. Clone the repository:

```bash
git clone <your-repo-url>
cd <repo-directory>
```

2. Initialize Go module (if not already done):

```bash
go mod init hello-world-api
```

## Running the Application

Start the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

Expected output:

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

- `name` (query parameter, required): The name to validate

#### Response Scenarios

##### 1. Valid Name (First Letter A-M)

**Request:**

```bash
curl "http://localhost:8080/hello-world?name=Alice"
```

**Response:**

```json
{
  "message": "Hello Alice"
}
```

**Status Code:** `200 OK`

##### 2. Invalid Name (First Letter N-Z)

**Request:**

```bash
curl "http://localhost:8080/hello-world?name=Zane"
```

**Response:**

```json
{
  "error": "Invalid Input"
}
```

**Status Code:** `400 Bad Request`

##### 3. Missing/Empty Name

**Request:**

```bash
curl "http://localhost:8080/hello-world?name="
```

**Response:**

```json
{
  "error": "Invalid Input"
}
```

**Status Code:** `400 Bad Request`

## Design Decisions & Assumptions

### Assumptions Made

1. **Case Insensitivity**: Both uppercase and lowercase letters are handled (e.g., "Alice" and "alice" are both valid)
2. **Whitespace Handling**: Leading and trailing whitespace is trimmed from the name parameter
3. **First Half Definition**: A-M (inclusive) represents the first half of the alphabet
4. **Non-Letter Characters**: Names starting with numbers or special characters are treated as invalid
5. **HTTP Methods**: Only GET requests are supported; other methods return 405 Method Not Allowed
6. **Port Configuration**: The service runs on port 8080 by default

### Architecture & Code Structure

1. **Separation of Concerns**:
   - `isFirstHalfAlphabet()` handles the business logic for validation
   - `helloWorldHandler()` manages HTTP request/response handling

2. **Error Handling**: All edge cases are handled gracefully with appropriate HTTP status codes

3. **Testing Strategy**:
   - Comprehensive unit tests covering happy paths, error cases, and edge cases
   - Table-driven tests for maintainability
   - Separate tests for helper functions

4. **JSON Responses**: Structured response types (`Response` and `ErrorResponse`) for consistency

5. **Standard Library Only**: Uses only Go's standard library (`net/http`, `encoding/json`) for minimal dependencies

## Project Structure

```
.
├── main.go           # Application entry point and HTTP handler
├── main_test.go      # Comprehensive unit tests
├── README.md         # This file
└── go.mod            # Go module file (if created)
```

## Testing Coverage

The test suite includes:

- Valid names (A-M, both cases)
- Invalid names (N-Z, both cases)
- Edge cases (empty strings, whitespace, single characters)
- Special characters and numbers as first character
- HTTP method validation
- Response format validation

## Error Handling

The service handles the following error scenarios:

- Missing `name` parameter
- Empty `name` parameter
- Whitespace-only `name` parameter
- Names starting with letters N-Z
- Names starting with non-letter characters
- Unsupported HTTP methods

## Production Considerations

For production deployment, consider:

- Adding configuration management for port and other settings
- Implementing structured logging
- Adding metrics and health check endpoints
- Implementing rate limiting
- Adding middleware for request logging
- Setting up graceful shutdown
- Containerization with Docker

## Example Usage

```bash
# Valid requests
curl "http://localhost:8080/hello-world?name=Alice"    # Returns 200
curl "http://localhost:8080/hello-world?name=Bob"      # Returns 200
curl "http://localhost:8080/hello-world?name=Mary"     # Returns 200

# Invalid requests
curl "http://localhost:8080/hello-world?name=Nancy"    # Returns 400
curl "http://localhost:8080/hello-world?name=Zane"     # Returns 400
curl "http://localhost:8080/hello-world?name="         # Returns 400
curl "http://localhost:8080/hello-world"               # Returns 400
```

## License

This is a coding assessment project.

## Contact

For questions or issues, please open an issue in the repository.

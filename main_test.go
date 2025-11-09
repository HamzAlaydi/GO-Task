package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   any
	}{
		{
			name:           "valid name starting with A",
			queryParam:     "Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello Alice"},
		},
		{
			name:           "valid name starting with M",
			queryParam:     "Mary",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello Mary"},
		},
		{
			name:           "valid lowercase name",
			queryParam:     "bob",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello bob"},
		},
		{
			name:           "invalid name starting with N",
			queryParam:     "Nancy",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
		{
			name:           "invalid name starting with Z",
			queryParam:     "Zane",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
		{
			name:           "invalid lowercase name second half",
			queryParam:     "oscar",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
		{
			name:           "empty name parameter",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
		{
			name:           "whitespace name parameter",
			queryParam:     "   ",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
		{
			name:           "name with spaces trimmed",
			queryParam:     "  Alice  ",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello Alice"},
		},
		{
			name:           "name starting with number",
			queryParam:     "123John",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
		{
			name:           "single character valid",
			queryParam:     "A",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello A"},
		},
		{
			name:           "single character invalid",
			queryParam:     "Z",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorResponse{Error: errorInvalidInput},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := newHelloWorldRequest(http.MethodGet, tt.queryParam)
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(helloWorldHandler)
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
				t.Fatalf("expected Content-Type application/json, got %s", contentType)
			}

			switch expected := tt.expectedBody.(type) {
			case Response:
				var actual Response
				if err := json.NewDecoder(rec.Body).Decode(&actual); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if actual != expected {
					t.Fatalf("expected body %#v, got %#v", expected, actual)
				}
			case ErrorResponse:
				var actual ErrorResponse
				if err := json.NewDecoder(rec.Body).Decode(&actual); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}
				if actual != expected {
					t.Fatalf("expected error %#v, got %#v", expected, actual)
				}
			default:
				t.Fatalf("unexpected expectedBody type %T", expected)
			}
		})
	}
}

func TestHelloWorldHandler_MethodNotAllowed(t *testing.T) {
	methods := []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := newHelloWorldRequest(method, "Alice")
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(helloWorldHandler)
			handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusMethodNotAllowed {
				t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
			}

			var body ErrorResponse
			if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
				t.Fatalf("failed to decode error response: %v", err)
			}

			if body.Error != errorMethodNotAllowed {
				t.Fatalf("expected error %q, got %q", errorMethodNotAllowed, body.Error)
			}
		})
	}
}

func TestIsFirstHalfAlphabet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"uppercase A", "Alice", true},
		{"uppercase M", "Mary", true},
		{"lowercase a", "alice", true},
		{"lowercase m", "mary", true},
		{"uppercase N", "Nancy", false},
		{"uppercase Z", "Zane", false},
		{"lowercase n", "nancy", false},
		{"lowercase z", "zane", false},
		{"empty string", "", false},
		{"number first", "123", false},
		{"special char first", "@name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isFirstHalfAlphabet(tt.input)
			if result != tt.expected {
				t.Fatalf("isFirstHalfAlphabet(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func newHelloWorldRequest(method, name string) *http.Request {
	req := httptest.NewRequest(method, "/hello-world", nil)
	q := req.URL.Query()
	q.Set("name", name)
	req.URL.RawQuery = q.Encode()
	return req
}


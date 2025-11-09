package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	successPort          = ":8080"
	errorInvalidInput    = "Invalid Input"
	errorMethodNotAllowed = "Method not allowed"
)

// Response represents the success response structure.
type Response struct {
	Message string `json:"message"`
}

// ErrorResponse represents the error response structure.
type ErrorResponse struct {
	Error string `json:"error"`
}

// isFirstHalfAlphabet returns true when the first character of name is a letter
// that falls within the first half of the English alphabet (A–M).
func isFirstHalfAlphabet(name string) bool {
	if name == "" {
		return false
	}

	first, _ := utf8.DecodeRuneInString(name)
	if first == utf8.RuneError {
		return false
	}

	if !unicode.IsLetter(first) {
		return false
	}

	first = unicode.ToUpper(first)
	return first >= 'A' && first <= 'M'
}

// helloWorldHandler handles the /hello-world endpoint.
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: errorMethodNotAllowed})
		return
	}

	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" || !isFirstHalfAlphabet(name) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: errorInvalidInput})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response{Message: "Hello " + name})
}

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)

	log.Printf("Server starting on port %s", successPort)

	if err := http.ListenAndServe(successPort, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}


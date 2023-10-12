package auth

import (
	"errors"
	"net/http"
	"strings"
)

// TODO: Example (check how this integrates with Gin)
func ExtractAPIAccessToken(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header")
	}

	values := strings.Split(val, " ")
	if len(values) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if values[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return values[1], nil
}

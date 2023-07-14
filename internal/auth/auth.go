package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers
// Example
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization header provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of Authorization header")
	}

	return vals[1], nil
}
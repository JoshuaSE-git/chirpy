package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	headerValue := headers.Get("Authorization")
	if headerValue == "" {
		return "", errors.New("no api key provided")
	}
	words := strings.Fields(headerValue)
	return words[1], nil
}

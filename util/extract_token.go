package util

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractTokenBearerFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	authorizations := strings.Split(authHeader, " ")
	if len(authorizations) < 1 {
		return "", errors.New("malformed authorization header")
	}

	if authorizations[0] != "Bearer" {
		return "", errors.New("authorization header must be 'Bearer'")
	}

	if len(authorizations[1]) == 0 {
		return "", errors.New("token is missing")
	}

	return authorizations[1], nil
}

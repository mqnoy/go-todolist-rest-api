package domain

import (
	"net/http"
)

type MiddlewareAuthorization interface {
	AuthorizationJWT(h http.Handler) http.Handler
}

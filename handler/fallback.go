package handler

import (
	"fmt"
	"net/http"
)

func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	ParseToErrorMsg(w, r, http.StatusNotFound, fmt.Errorf("%s not found", r.RequestURI))
}

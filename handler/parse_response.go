package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mqnoy/go-todolist-rest-api/pkg/cerror"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DefaultResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ParseToErrorMsg(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	render.JSON(w, r, DefaultResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}

func ParseToDefaultMessage(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	render.JSON(w, r, DefaultResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ParseResponse(w http.ResponseWriter, r *http.Request, message string, data interface{}, err error) {
	var customErr *cerror.CustomError

	if err != nil {
		if errors.As(err, &customErr) {
			ParseToErrorMsg(w, r, customErr.StatusCode, customErr.Err)
			return
		}

		ParseToErrorMsg(w, r, http.StatusInternalServerError, err)
		return
	}

	ParseToDefaultMessage(w, r, message, data)
}

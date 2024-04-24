package dto

import (
	"net/http"

	"gopkg.in/guregu/null.v4"
)

type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     int64    `json:"dueDate"`
	DoneAt      null.Int `json:"doneAt"`
	Timestamp
}

type TaskCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueDate     int64  `json:"dueDate" validate:"required"`
}

func (t *TaskCreateRequest) Bind(r *http.Request) error {
	return nil
}

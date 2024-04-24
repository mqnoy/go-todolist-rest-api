package dto

import "gopkg.in/guregu/null.v4"

type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     int64    `json:"dueDate"`
	DoneAt      null.Int `json:"doneAt"`
}

package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Task struct {
	UUIDColumn
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	DueDate     time.Time `gorm:"column:dueDate"`
	DoneAt      null.Time `gorm:"column:isDone"`
	TimestampColumn
}

func (m *Task) TableName() string {
	return "Task"
}
